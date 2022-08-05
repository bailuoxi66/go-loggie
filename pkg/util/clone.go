package util

import "gopkg.in/yaml.v2"

func Clone(a, b interface{}) error {
	out, err := yaml.Marshal(a)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(out, b)
}
