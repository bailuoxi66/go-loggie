package queue

import "bailuoxi66/go-loggie/pkg/core/cfg"

type Config struct {
	cfg.ComponentBaseConfig `yaml:",inline"`
	BatchSize               int `yaml:"batchSize,omitempty" default:"2048" validate:"required,gte=1"`
}
