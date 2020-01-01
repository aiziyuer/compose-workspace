package util

import (
	"fmt"
	"net/http"
	"regexp"
)

func WrapperRequest(s string) (*http.Request, error) {

	m, err := JsonX2Map(s)
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("%s://%s", m["Schema"].(string), m["Host"].(string))
	realUrl := Url(prefix, m["Path"].(string))

	r, err := http.NewRequest(
		m["Method"].(string),
		realUrl,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return r, nil

}

func Url(prefix string, uri string, args ...interface{}) string {

	tmpUrl := fmt.Sprintf("%s/%s", prefix, fmt.Sprintf(uri, args...))

	m := RegexNamedMatch(tmpUrl, `(?P<schema>^\w+://|^)(?P<host>[^/]+)(?P<path>[\w/]+)`)
	schema := m["schema"]
	if schema == "" {
		schema = "https://"
	}
	host := m["host"]
	path := m["path"]

	tmpSuffix := fmt.Sprintf("%s/%s", host, path)
	suffix := regexp.MustCompile(`[/]+`).ReplaceAllString(tmpSuffix, `/`)
	url := fmt.Sprintf("%s%s", schema, suffix)
	return url
}

func RegexNamedMatch(input string, pattern string) map[string]string {

	re := regexp.MustCompile(pattern)
	m := map[string]string{}
	for i, v := range re.FindStringSubmatch(input) {
		name := re.SubexpNames()[i]
		if name == "" {
			continue
		}
		m[name] = v
	}

	return m
}
