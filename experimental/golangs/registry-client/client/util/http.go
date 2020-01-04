package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

func Url(args ...string) string {
	return UrlWithSeparator("", args...)
}

func UrlWithSeparator(sep string, args ...string) string {

	tmpUrl := strings.Join(args, sep)

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

func ReadWithDefault(r io.Reader, fallback string) string {

	ret, err := ioutil.ReadAll(r)
	if err != nil {

		return fallback
	}

	return string(ret)
}
