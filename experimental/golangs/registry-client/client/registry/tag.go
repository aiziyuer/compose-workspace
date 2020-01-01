package registry

import (
	"fmt"
	"github.com/aiziyuer/registry/client/util"
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
	out, _ := util.TemplateRenderByPong2(request, context)
	req, _ := util.WrapperRequest(out)
	res, _ := r.Handler.DoWithContext(req, &context)

	fmt.Println(res)
	return nil, nil
}
