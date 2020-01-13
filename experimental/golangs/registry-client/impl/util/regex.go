package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"regexp"
)

func RegexNamedMatch(input string, pattern string) map[string]string {

	m := map[string]string{}

	re, err := regexp.Compile(pattern)
	if err != nil {
		logrus.Warn(fmt.Sprintf("input(%s), pattern(%s)", input, pattern), err)
		return m
	}

	for i, v := range re.FindStringSubmatch(input) {
		name := re.SubexpNames()[i]
		if name == "" {
			continue
		}
		m[name] = v
	}

	return m
}
