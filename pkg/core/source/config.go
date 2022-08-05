package source

import (
	"bailuoxi66/go-loggie/pkg/core/cfg"
	"github.com/pkg/errors"
)

var (
	ErrSourceNameRequired = errors.New("source name is required")
)

type Config struct {
	cfg.ComponentBaseConfig `yaml:",inline"`
	FieldsUnderRoot         bool                   `yaml:"fieldsUnderRoot,omitempty" default:"false"`
	FieldsUnderKey          string                 `yaml:"fieldsUnderKey,omitempty" default:"fields"`
	Fields                  map[string]interface{} `yaml:"fields,omitempty"`
}

func (c *Config) Validate() error {
	if c.Name == "" {
		return ErrSourceNameRequired
	}
	return nil
}
