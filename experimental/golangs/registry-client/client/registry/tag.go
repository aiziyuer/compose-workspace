package registry

import (
	"github.com/aiziyuer/registry/client/handler"
	"github.com/aiziyuer/registry/client/util"
	"github.com/sirupsen/logrus"
)

func (r *Registry) Tags(repoName string) (string, error) {

	q, err := handler.NewApiRequest(handler.ApiRequestInput{
		"Schema":   r.Endpoint.Schema,
		"Host":     r.Endpoint.Host,
		"RepoName": repoName,
		"Token":    "",
	}, `
	{
		"Method": "GET",
		"Path": "/v2/{{ .RepoName }}/tags/list",
		"Schema": "{{ .Schema }}",
		"Host": "{{ .Host }}",
		"Header": {
			"Content-Type": "application/json; charset=utf-8",
			"Authorization": "{{ .Token }}",
		},
		"Body": "",
	}
`)
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
