package sink

import (
	"bailuoxi66/go-loggie/pkg/core/cfg"
	"bailuoxi66/go-loggie/pkg/sink/codec"
)

type Config struct {
	cfg.ComponentBaseConfig `yaml:",inline"`
	Parallelism             int          `yaml:"parallelism,omitempty" default:"1" validate:"required,gte=1,lte=100"`
	Codec                   codec.Config `yaml:"codec" validate:"dive"`
}

func (c *Config) Validate() error {
	return c.Codec.Validate()
}
