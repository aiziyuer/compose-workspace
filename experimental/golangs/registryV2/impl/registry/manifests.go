package registry

import (
	"errors"
	"fmt"
	"github.com/aiziyuer/registryV2/impl/handler"
	"github.com/aiziyuer/registryV2/impl/util"
	"github.com/sirupsen/logrus"
)

func (r *Registry) Manifests(imageName string) (string, error) {

	m := util.RegexNamedMatch(imageName, `(?P<RepoName>[a-z0-9]+(?:[._\-/][a-z0-9]+)*):(?P<TagName>[a-z0-9]+(?:[._\-/][a-z0-9]+)*)`)
	if len(m) == 0 {
		return "", errors.New(fmt.Sprintf("image name(%s) invalid", imageName))
	}

	q, err := handler.NewApiRequest(`
{
    "Method": "GET",
    "Path": "/v2/{{ .RepoName }}/manifests/{{ .TagName }}",
    "Schema": "{{ .Schema }}",
    "Host": "{{ .Host }}",
    "Headers": {
        "Accept-Encoding": "gzip",
        "Content-Type": "application/json; charset=utf-8",
        "User-Agent": "docker/1.13.1 go/go1.10.3 kernel/3.10.0-1062.4.1.el7.x86_64 os/linux arch/amd64 UpstreamClient(Docker-Client/1.13.1 \\(linux\\))",
        "Accept": [
            "application/json",
            "application/vnd.docker.distribution.manifest.v2+json",
            "application/vnd.docker.distribution.manifest.list.v2+json",
            "application/vnd.docker.distribution.manifest.v1+prettyjws"
        ],
        "Authorization": "{{ .Token }}"
    },
    "Body": ""
}
`, handler.ApiRequestInput{
		"Schema":   r.Endpoint.Schema,
		"Host":     r.Endpoint.Host,
		"RepoName": m["RepoName"],
		"TagName":  m["TagName"],
		"Token":    "",
	})
	if err != nil {
		return "", err
	}

	req, _ := q.Wrapper()
	res, err := r.HandlerFacade.Do(req)
	if err != nil {
		return "", err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			logrus.Errorf("res.Body.Close() error: ", err)
		}
	}()

	return util.ReadWithDefault(res.Body, "{}"), nil
}
