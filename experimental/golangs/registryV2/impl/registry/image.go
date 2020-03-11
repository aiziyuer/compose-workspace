package registry

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/aiziyuer/registryV2/impl/handler"
	"github.com/aiziyuer/registryV2/impl/util"
	"github.com/pkg/math"
	"github.com/sirupsen/logrus"
)

type (
	Project struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		StartCount  int    `json:"star_count"`
		Trusted     bool   `json:"is_trusted"`
		Automated   bool   `json:"is_automated"`
		Official    bool   `json:"is_official"`
	}

	ProjectSearchResult struct {
		PagesTotalNum int       `json:"num_pages"`
		TotalSize     int       `json:"num_results"`
		PageSize      int       `json:"page_size"`
		PageNo        int       `json:"page"`
		Projects      []Project `json:"results"`
	}

	Platform struct {
		Architecture string `json:"architecture"`
		OS           string `json:"os"`
	}

	ManifestConfig struct {
		Architecture string `json:"architecture"`
		OS           string `json:"os"`
		MediaType    string `json:"mediaType"`
		Size         int    `json:"size"`
		Digest       string `json:"digest"`
	}

	MediaType struct {
		MediaType string `json:"mediaType"`
		Size      int    `json:"size"`
		Digest    string `json:"digest"`
	}

	SubManifestsV2 struct {
		Digest        string      `json:"digest"`
		MediaType     string      `json:"mediaType"`
		Config        *MediaType  `json:"config"`
		Layers        []MediaType `json:"layers"`
		Platform      *Platform   `json:"platform"`
		Size          int         `json:"size"`
		SchemaVersion int         `json:"schemaVersion"`
		RawBase64     string      `json:"rawBase64"`
		RawSha256Sum  string      `json:"rawSha256Sum"`
	}

	ManifestV2 struct {
		Digest        string           `json:"digest"`
		SchemaVersion int              `json:"schemaVersion"`
		MediaType     string           `json:"mediaType"`
		Manifests     []SubManifestsV2 `json:"manifests"`
		Size          int              `json:"size"`
		RawBase64     string           `json:"rawBase64"`
		RawSha256Sum  string           `json:"rawSha256Sum"`
	}
)

const SearchProjectRequestTemplate = `
{
	"Method": "GET",
	"Path": "/v1/search",
	"Schema": "{{ .Schema }}",
	"Host": "{{ .Host }}",
	"Headers": {
		"Content-Type": "application/json; charset=utf-8",
		"Authorization": "{{ .Authorization }}",
	},
	"Params": {
		"q": "{{ .NameQuery }}",
		"n": "{{ .PageSize }}",
		"page": "{{ .PageNo }}",
	},
	"Body": "",
}
`

func (r *Registry) searchProject(requestInput *handler.ApiRequestInput) (*ProjectSearchResult, error) {

	if r.Auth != nil {
		basicAuth := fmt.Sprintf("%s:%s", r.Auth.UserName, r.Auth.PassWord)
		encoded := base64.StdEncoding.EncodeToString([]byte(basicAuth))
		(*requestInput)["Authorization"] = fmt.Sprintf("Basic %s", encoded)
	}

	var result ProjectSearchResult

	if err := r.Do( //
		SearchProjectRequestTemplate, //
		requestInput,                 //
		func(resp *http.Response) error {

			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				logrus.Error(err)
				return err
			}

			return nil
		}, //
	); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &result, nil
}

func (r *Registry) SearchProject(nameQuery string, n int) ([]Project, error) {

	var projects = make([]Project, 0)

	ch := make(chan []Project)
	defer close(ch)
	var jobWg, dataWg sync.WaitGroup

	// 尝试获取服务器的总结果
	tmpResp, err := r.searchProject(&handler.ApiRequestInput{
		"Schema":    r.Endpoint.Schema,
		"Host":      "index.docker.io",
		"NameQuery": nameQuery,
		"PageNo":    1,
		"PageSize":  25,
	})
	if err != nil {
		return nil, err
	}

	// 请求所有
	if n == -1 {
		n = tmpResp.TotalSize
	}
	n = math.Min(n, tmpResp.TotalSize)

	// 消费协程
	go func() {
		for project := range ch {
			projects = append(projects, project...)
			dataWg.Done() // 务必保证数据消费掉才结束
		}
	}()

	// 生产协程
	for i := 0; i < n/tmpResp.PageSize+1; i++ {

		jobWg.Add(1)

		go func(pageNo int) {

			defer jobWg.Done()

			result, err := r.searchProject(&handler.ApiRequestInput{
				"Schema":    r.Endpoint.Schema,
				"Host":      "index.docker.io",
				"NameQuery": nameQuery,
				"PageSize":  tmpResp.PageSize,
				"PageNo":    pageNo,
			})
			if err != nil {
				return
			}

			dataWg.Add(1)
			ch <- result.Projects

		}(i + 1) // 这里形参传递是防止goroutine的执行不可控, 里面对变量i的访问不好预测

	}

	// 等待任务结束
	jobWg.Wait()
	// 等待数据处理完
	dataWg.Wait()

	// 按照start数进行排序
	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].StartCount > projects[j].StartCount
	})

	return projects, nil
}

const V2ManifestRequestTemplate = `
{
    "Method": "GET",
    "Path": "/v2/{{ .RepoName}}/manifests/{{ .Index }}",
    "Schema": "{{ .Schema }}",
    "Host": "{{ .Host }}",
    "Headers": {
        "Accept-Encoding": "gzip",
        "Content-Type": "application/json; charset=utf-8",
        "Authorization": "",
        "User-Agent": "docker/1.13.1 go/go1.10.3 kernel/3.10.0-1062.4.1.el7.x86_64 os/linux arch/amd64 UpstreamClient(Docker-Client/1.13.1 \\(linux\\))",
        "Accept": [
            "application/vnd.docker.distribution.manifest.v2+json",
            "application/vnd.docker.distribution.manifest.list.v2+json",
            "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "application/json"
        ]
    }
}
`

const V2BlobRequestTemplate = `
{
    "Method": "GET",
    "Path": "/v2/{{ .RepoName}}/blobs/{{ .Index }}",
    "Schema": "{{ .Schema }}",
    "Host": "{{ .Host }}",
    "Headers": {
        "Accept-Encoding": "identity",
        "Content-Type": "application/json; charset=utf-8",
        "Authorization": "",
        "User-Agent": "docker/1.13.1 go/go1.10.3 kernel/3.10.0-1062.4.1.el7.x86_64 os/linux arch/amd64 UpstreamClient(Docker-Client/1.13.1 \\(linux\\))"
    }
}
`

func (r *Registry) ManifestV2(imageFullName string) (*ManifestV2, error) {

	manifestV2 := &ManifestV2{
		Size: 0,
		Manifests: []SubManifestsV2{
			{
				Config: nil,
				Size:   0,
			},
		},
	}

	m := util.RegexNamedMatch(imageFullName, `(?:(?P<Host>^[^.]+\.[^/]+)/)?(?P<RepoName>[a-z0-9]+(?:[._\-/][a-z0-9]+)*):(?P<TagName>[a-z0-9]+(?:[._\-/][a-z0-9]+)*)`)
	if len(m) == 0 {
		return manifestV2, errors.New(fmt.Sprintf("image name(%s) invalid", imageFullName))
	}

	host := r.Endpoint.Host
	if m["Host"] != "" {
		host = m["Host"]
	}

	repoName := m["RepoName"]
	if m["Host"] == "" && !strings.HasPrefix(m["RepoName"], "library") {
		repoName = fmt.Sprintf("library/%s", m["RepoName"])
	}

	ch := make(chan *SubManifestsV2)
	defer close(ch)
	var jobWg, dataWg sync.WaitGroup

	// 消费协程
	go func() {

		for subManifestsV2 := range ch {

			func() {

				defer dataWg.Done()

				var tmpBody string
				if err := r.Do(
					V2ManifestRequestTemplate,
					&handler.ApiRequestInput{
						"Schema":   r.Endpoint.Schema,
						"Host":     host,
						"RepoName": repoName,
						"Index":    subManifestsV2.Digest,
					}, //
					func(resp *http.Response) error {

						subManifestsV2.Digest = resp.Header.Get("Docker-Content-Digest")
						subManifestsV2.RawBase64 = base64.RawStdEncoding.EncodeToString([]byte(tmpBody))

						tmpBody = util.ReadWithDefault(resp.Body, "{}")
						h := sha256.New()
						h.Write([]byte(tmpBody))
						subManifestsV2.RawSha256Sum = hex.EncodeToString(h.Sum(nil))

						return nil
					},
				); err != nil {
					logrus.Error(err)
					return
				}

				if err := json.Unmarshal([]byte(tmpBody), subManifestsV2); err != nil {
					logrus.Error(err)
				}

				subManifestsV2.Size = subManifestsV2.Config.Size
				for _, layer := range subManifestsV2.Layers {
					subManifestsV2.Size += layer.Size
				}

				if err := r.Do(
					V2BlobRequestTemplate,
					&handler.ApiRequestInput{
						"Schema":   r.Endpoint.Schema,
						"Host":     host,
						"RepoName": repoName,
						"Index":    subManifestsV2.Config.Digest,
					}, //
					func(resp *http.Response) error {

						config := ManifestConfig{}
						if err := json.NewDecoder(resp.Body).Decode(&config); err != nil {
							return err
						}

						subManifestsV2.Platform = &Platform{
							Architecture: config.Architecture,
							OS:           config.OS,
						}
						return nil
					}, //
				); err != nil {
					logrus.Error(err)
					return
				}

			}()
		}
	}()

	// 生产协程
	jobWg.Add(1)
	go func() {
		defer jobWg.Done()

		var tmpBody string
		if err := r.Do(
			V2ManifestRequestTemplate,
			&handler.ApiRequestInput{
				"Schema":   r.Endpoint.Schema,
				"Host":     host,
				"RepoName": repoName,
				"Index":    m["TagName"],
			}, //
			func(resp *http.Response) error {

				manifestV2.Digest = resp.Header.Get("Docker-Content-Digest")
				manifestV2.RawBase64 = base64.RawStdEncoding.EncodeToString([]byte(tmpBody))

				tmpBody = util.ReadWithDefault(resp.Body, "{}")
				h := sha256.New()
				h.Write([]byte(tmpBody))
				manifestV2.RawSha256Sum = hex.EncodeToString(h.Sum(nil))

				return nil
			},
		); err != nil {
			logrus.Error(err)
			return
		}

		if strings.Contains(tmpBody, "manifests") {

			if err := util.Json2Object(tmpBody, manifestV2); err != nil {
				logrus.Error(err)
				return
			}

			for i := 0; i < len(manifestV2.Manifests); i++ {
				dataWg.Add(1)
				ch <- &manifestV2.Manifests[i]
			}

		} else {

			manifestV2.Manifests[0].Digest = manifestV2.Digest
			dataWg.Add(1)
			ch <- &manifestV2.Manifests[0]

		}
	}()

	// 等待任务结束
	jobWg.Wait()
	// 等待数据处理完
	dataWg.Wait()

	manifestV2.Size = 0
	for _, manifest := range manifestV2.Manifests {
		manifestV2.Size += manifest.Size
	}

	return manifestV2, nil
}
