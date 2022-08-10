package transformer

import (
	"bailuoxi66/go-loggie/pkg/util"
	"github.com/bitly/go-simplejson"
)

type RenameProcessor struct {
	config *RenameConfig
}

type RenameConfig struct {
	Target []Convert `yaml:"target,omitempty"`
}

func init() {
	register("rename", func() Processor {
		return NewRenameProcessor()
	})
}

func NewRenameProcessor() *RenameProcessor {
	return &RenameProcessor{
		config: &RenameConfig{},
	}
}

func (d *RenameProcessor) Config() interface{} {
	return d.config
}

func (d *RenameProcessor) Process(jsonObj *simplejson.Json) {
	if d.config == nil {
		return
	}
	for _, t := range d.config.Target {
		from := t.From
		to := t.To

		fromPaths := util.GetQueryPaths(from)
		toPaths := util.GetQueryPaths(to)
		tmp := jsonObj.GetPath(fromPaths...)
		jsonObj.SetPath(toPaths, tmp.Interface())

		upperFromPaths, key := util.GetQueryUpperPaths(from)
		upperFrom := jsonObj.GetPath(upperFromPaths...)
		upperFrom.Del(key)
	}
}
