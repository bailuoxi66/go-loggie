package transformer

import (
	"bailuoxi66/go-loggie/pkg/core/api"
	"bailuoxi66/go-loggie/pkg/core/cfg"
	"bailuoxi66/go-loggie/pkg/core/log"
	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
)

type Transformer struct {
	processor []Processor
}

func NewTransformer(config TransConfig) *Transformer {
	if len(config) == 0 {
		return &Transformer{}
	}

	var processors []Processor
	for _, conf := range config {
		for name, properties := range conf {
			proc, err := newProcessor(name, properties)
			if err != nil {
				log.Warn("get processor error: %+v", err)
				continue
			}
			processors = append(processors, proc)
		}
	}

	return &Transformer{
		processor: processors,
	}
}

func newProcessor(name string, properties cfg.CommonCfg) (Processor, error) {
	proc, ok := getProcessor(name)
	if !ok {
		return nil, errors.Errorf("transformer %s cannot be found", name)
	}
	if c, ok := proc.(api.Config); ok {
		err := cfg.UnpackAndDefaults(properties, c.Config())
		if err != nil {
			return nil, errors.WithMessagef(err, "unpack transformer %s config", name)
		}
	}
	return proc, nil
}

func (t *Transformer) Transform(jsonObj *simplejson.Json) {
	for _, p := range t.processor {
		p.Process(jsonObj)
	}
}
