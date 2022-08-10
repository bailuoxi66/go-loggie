package codec

import (
	"bailuoxi66/go-loggie/pkg/core/cfg"
	"github.com/pkg/errors"
)

type Config struct {
	Type          string `yaml:"type" default:"json"`
	cfg.CommonCfg `yaml:",inline"`
}

func (c *Config) Validate() error {
	if c.Type != "json" {
		return errors.Errorf("codec %s is not supported", c.Type)
	}
	return nil
}
