package registry

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aiziyuer/registryV2/impl/handler"
	"github.com/pkg/math"
	"sort"
	"sync"
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
)

func (r *Registry) searchProject(requestInput *handler.ApiRequestInput) (*ProjectSearchResult, error) {

	if r.Auth != nil {
		basicAuth := fmt.Sprintf("%s:%s", r.Auth.UserName, r.Auth.PassWord)
		encoded := base64.StdEncoding.EncodeToString([]byte(basicAuth))
		(*requestInput)["Authorization"] = fmt.Sprintf("Basic %s", encoded)
	}

	q, err := handler.NewApiRequest(*requestInput, `
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
`)
	if err != nil {
		return nil, err
	}

	req, _ := q.Wrapper()
	resp, _ := r.HandlerFacade.Do(req)
	if resp != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}

	var result ProjectSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
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
