package registry

import (
	"fmt"
)

func (r *Registry) Tags(repoName string) (tags []string, err error) {

	context := map[string]interface{}{
		"RepoName": repoName,
	}

	request := `
{
    // this json syntax is jsonx: https://github.com/danharper/JSONx 
    // and variable render syntax like django
    "Method": "GET",
    "Path": "/v2/{{ RepoName }}/tags/list",
    "Schema": "https",
    "Host": "registry-1.docker.io",
    "Header": {
        "Token": "{{ Token | default: "123" }}",
    },
    "Body": "",
}
`
	res, _ := r.Handler.DoWithContextV2(request, &context)

	fmt.Println(res)
	return nil, nil
}
