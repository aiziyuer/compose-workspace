package registry

import (
	"github.com/aiziyuer/registryV2/impl/handler"
	"github.com/aiziyuer/registryV2/impl/util"
	"github.com/sirupsen/logrus"
)

func (r *Registry) Tags(repoName string) (string, error) {

	q, err := handler.NewApiRequest(`
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
`, handler.ApiRequestInput{
		"Schema":   r.Endpoint.Schema,
		"Host":     r.Endpoint.Host,
		"RepoName": repoName,
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

func (r *Registry) TagsPaginated(repoName string, pageNo int, pageSize int) (string, error) {

	q, err := handler.NewApiRequest(`
	{
		"Method": "GET",
		"Path": "/v2/{{ .RepoName }}/tags/list",
        "Params": {
            "last": "{{ .PageNo }}",
            "n": "{{ .PageSize }}",
        },
		"Schema": "{{ .Schema }}",
		"Host": "{{ .Host }}",
		"Header": {
			"Content-Type": "application/json; charset=utf-8",
			"Authorization": "{{ .Token }}",
		},
		"Body": "",
	}
`, handler.ApiRequestInput{
		"Schema":   r.Endpoint.Schema,
		"Host":     r.Endpoint.Host,
		"RepoName": repoName,
		"PageNo":   pageNo,   // Result set will include values lexically after last.
		"PageSize": pageSize, // Limit the number of entries in each response. It not present, all entries will be returned.
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
