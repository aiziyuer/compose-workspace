package registry

import (
	"github.com/aiziyuer/registry/client/handler"
	"github.com/hokaccha/go-prettyjson"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func (r *Registry) TagsWithPretty(repoName string) (string, error) {
	originalJson, err := r.Tags(repoName)
	if err != nil {
		return "", nil
	}

	s, err := prettyjson.Format([]byte(originalJson))
	if err != nil {
		return "", nil
	}

	return string(s), nil
}

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
	res, _ := r.Handler.Do(req)
	defer func() {
		if err := res.Body.Close(); err != nil {
			logrus.Errorf("res.Body.Close() error: ", err)
		}
	}()

	ret, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", nil
	}

	return string(ret), nil
}
