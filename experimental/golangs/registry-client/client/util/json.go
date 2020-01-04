package util

import (
	"github.com/hokaccha/go-prettyjson"
	"github.com/mkideal/pkg/encoding/jsonx"
	"github.com/sirupsen/logrus"
)

func JsonX2Map(s string) (map[string]interface{}, error) {

	var ret map[string]interface{}
	err := jsonx.Unmarshal([]byte(s), &ret, jsonx.WithExtraComma(), jsonx.WithComment())
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func JsonX2Object(s string, v interface{}) error {

	err := jsonx.Unmarshal([]byte(s), &v, jsonx.WithExtraComma(), jsonx.WithComment())
	if err != nil {
		return err
	}

	return nil
}

func PrettyFormat(s string) string {

	ret, err := prettyjson.Format([]byte(s))
	if err != nil {
		logrus.Warn("PrettyFormat with error:", err)
		return ""
	}

	return string(ret)
}
