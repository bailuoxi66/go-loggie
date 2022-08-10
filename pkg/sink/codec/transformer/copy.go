package transformer

import (
	"bailuoxi66/go-loggie/pkg/util"
	"github.com/bitly/go-simplejson"
)

type CopyProcessor struct {
	config *CopyConfig
}

type CopyConfig struct {
	Target []Convert `yaml:"target,omitempty"`
}

func init() {
	register("copy", func() Processor {
		return NewCopyProcessor()
	})
}

func NewCopyProcessor() *CopyProcessor {
	return &CopyProcessor{
		config: &CopyConfig{},
	}
}

func (d *CopyProcessor) Config() interface{} {
	return d.config
}

func (d *CopyProcessor) Process(jsonObj *simplejson.Json) {
	if d.config == nil {
		return
	}
	for _, t := range d.config.Target {
		from := t.From
		to := t.To
		fromPath := util.GetQueryPaths(from)
		toPath := util.GetQueryPaths(to)
		val := jsonObj.GetPath(fromPath...)
		jsonObj.SetPath(toPath, val)
	}
}
