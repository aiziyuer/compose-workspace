package util

import (
	"encoding/json"
	"github.com/hokaccha/go-prettyjson"
	"github.com/mkideal/pkg/encoding/jsonx"
	"github.com/sirupsen/logrus"
)

func JsonX2Object(s string, v interface{}) error {

	err := jsonx.Unmarshal([]byte(s), &v, jsonx.WithExtraComma(), jsonx.WithComment())
	if err != nil {
		return err
	}

	return nil
}

func Json2Object(s string, v interface{}) error {

	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		return err
	}

	return nil
}

func Object2Json(v interface{}) (string, error) {

	ret, err := jsonx.Marshal(v, jsonx.WithExtraComma(), jsonx.WithComment())
	if err != nil {
		return "", err
	}

	return string(ret), nil
}

func Object2PrettyJson(v interface{}) (string, error) {

	ret, err := json.Marshal(&v)
	if err != nil {
		return "", err
	}

	return string(PrettyJsonBytes(ret)), nil
}

func Object2JsonBytes(v interface{}) ([]byte, error) {

	ret, err := json.Marshal(&v)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func JsonBytesX2Object(s []byte, v interface{}) error {

	err := jsonx.Unmarshal(s, &v, jsonx.WithExtraComma(), jsonx.WithComment())
	if err != nil {
		return err
	}

	return nil
}

func PrettyJson(s string) string {

	ret, err := prettyjson.Format([]byte(s))
	if err != nil {
		logrus.Warn("PrettyJson with error:", err)
		return ""
	}

	return string(ret)
}

func PrettyJsonBytes(s []byte) []byte {

	ret, err := prettyjson.Format(s)
	if err != nil {
		logrus.Warn("PrettyJson with error:", err)
		return nil
	}

	return ret
}
