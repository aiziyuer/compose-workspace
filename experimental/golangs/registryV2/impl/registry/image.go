package registry

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aiziyuer/registryV2/impl/handler"
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

	ch := make(chan []Project, 10)
	defer close(ch)
	var wg sync.WaitGroup

	// 消费协程
	go func() {
		for project := range ch {
			projects = append(projects, project...)
		}
	}()

	// 尝试获取服务器的总结果
	result, err := r.searchProject(&handler.ApiRequestInput{
		"Schema":    r.Endpoint.Schema,
		"Host":      "index.docker.io",
		"NameQuery": nameQuery,
		"PageSize":  25,
	})
	if err != nil {
		return nil, err
	}

	// 请求所有
	if n == -1 {
		n = result.TotalSize
	}

	// 生产协程
	for i := 0; i < n/result.PageSize+1; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			result, err := r.searchProject(&handler.ApiRequestInput{
				"Schema":    r.Endpoint.Schema,
				"Host":      "index.docker.io",
				"NameQuery": nameQuery,
				"PageSize":  result.PageSize,
				"PageNo":    i + 1,
			})
			if err != nil {
				return
			}

			ch <- result.Projects
		}()

	}

	wg.Wait()

	// 按照start数进行排序
	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].StartCount > projects[j].StartCount
	})

	return projects, nil
}
