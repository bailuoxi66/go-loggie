package transformer

import (
	"bailuoxi66/go-loggie/pkg/util"
	"github.com/bitly/go-simplejson"
)

type AddProcessor struct {
	config *AddConfig
}

type AddConfig struct {
	Target map[string]interface{} `yaml:"target,omitempty"`
}

func init() {
	register("add", func() Processor {
		return NewAddProcessor()
	})
}

func NewAddProcessor() *AddProcessor {
	return &AddProcessor{
		config: &AddConfig{},
	}
}

func (d *AddProcessor) Config() interface{} {
	return d.config
}

func (d *AddProcessor) Process(jsonObj *simplejson.Json) {
	if d.config == nil {
		return
	}
	for k, v := range d.config.Target {
		paths := util.GetQueryPaths(k)
		jsonObj.SetPath(paths, v)
	}
}
