package registry

import (
	"fmt"
	"github.com/aiziyuer/registry/client/handler"
)

func (r *Registry) Tags(repoName string) (tags []string, err error) {

	q := (&handler.ApiRequest{
		Input: map[string]interface{}{
			"RepoName": repoName,
		},
		Template: `
{
    "Method": "GET",
    "Path": "/v2/{{ RepoName }}/tags/list",
    "Schema": "https",
    "Host": "registry-1.docker.io",
    "Header": {
        "Content-Type": "application/json",
    },
    "Body": "",
}
`,
	}).Render()

	req, _ := q.Wrapper()
	res, _ := r.Handler.Do(req)

	fmt.Println(res)
	return nil, nil
}
