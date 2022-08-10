package transformer

import "bailuoxi66/go-loggie/pkg/core/cfg"

type TransConfig []map[string]cfg.CommonCfg

type Convert struct {
	From string `yaml:"from,omitempty" validate:"required"`
	To   string `yaml:"to,omitempty" validate:"required"`
}
