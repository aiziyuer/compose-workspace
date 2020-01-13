package util

import (
	"bytes"
	"github.com/flosch/pongo2"
	"text/template"
)

func TemplateRender(src string, context interface{}) (string, error) {

	t := template.Must(template.New("").Parse(src))
	var dest bytes.Buffer
	err := t.Execute(&dest, context)
	if err != nil {
		return "", err
	}

	return dest.String(), err
}

func TemplateRenderByPong2(tpl string, context map[string]interface{}) (string, error) {

	var t = pongo2.Must(pongo2.DefaultSet.FromString(tpl))
	dst, err := t.Execute(context)
	if err != nil {
		return "", err
	}

	return dst, nil
}
