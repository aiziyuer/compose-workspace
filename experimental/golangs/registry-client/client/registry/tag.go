package registry

import (
	"github.com/aiziyuer/registry/client/util"
	"github.com/davecgh/go-spew/spew"
)

func (r *Registry) Tags(repoName string) (tags []string, err error) {

	context := map[string]interface{}{
		"RepoName": repoName,
	}

	request := `
{
    // django syntax
    "Schema": "https",
    "Host": "registry-1.docker.io",
    "Method": "GET",
    "Path": "/v2/{{ RepoName }}/tags/list",
    "Header": {
        "Token": "{{ Token | default:"123" }}",
    },
    "Body": "",
}
`
	out, _ := util.TemplateRenderByPong2(request, context)
	m, err := util.JsonX2Map(out)

	if err != nil {
		return nil, err
	}

	// TODO 反射获取Request中的数据进行渲染
	spew.Dump(m)

	return nil, nil
}
