package transformer

import (
	"bailuoxi66/go-loggie/pkg/util"
	"github.com/bitly/go-simplejson"
)

type DropProcessor struct {
	config *DropConfig
}

type DropConfig struct {
	Target []string `yaml:"target,omitempty"`
}

func init() {
	register("drop", func() Processor {
		return NewDropProcessor()
	})
}

func NewDropProcessor() *DropProcessor {
	return &DropProcessor{
		config: &DropConfig{},
	}
}

func (d *DropProcessor) Config() interface{} {
	return d.config
}

func (d *DropProcessor) Process(jsonObj *simplejson.Json) {
	if d.config == nil {
		return
	}
	for _, t := range d.config.Target {
		upperPath, lastQuery := util.GetQueryUpperPaths(t)
		upperObj := jsonObj.GetPath(upperPath...)
		upperObj.Del(lastQuery)
	}
}
