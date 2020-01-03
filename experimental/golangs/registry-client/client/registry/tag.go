package registry

import (
	"fmt"
	"github.com/aiziyuer/registry/client/handler"
)

func (r *Registry) Tags(repoName string) ([]string, error) {

	q, err := handler.NewApiRequest(handler.ApiRequestInput{
		"RepoName": repoName,
		"Token":    "",
	}, `
	{
		"Method": "GET",
		"Path": "/v2/{{ .RepoName }}/tags/list",
		"Schema": "https",
		"Host": "registry-1.docker.io",
		"Header": {
			"Content-Type": "application/json; charset=utf-8",
			"Authorization": "{{ .Token }}",
		},
		"Body": "",
	}
`)
	if err != nil {
		return nil, err
	}

	req, _ := q.Wrapper()
	res, _ := r.Handler.Do(req)

	fmt.Println(res)
	return nil, nil
}
