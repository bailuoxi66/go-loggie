package transformer

import (
	"bailuoxi66/go-loggie/pkg/util"
	"github.com/bitly/go-simplejson"
)

type MoveProcessor struct {
	config *MoveConfig
}

type MoveConfig struct {
	Target []Convert `yaml:"target,omitempty"`
}

func init() {
	register("move", func() Processor {
		return NewMoveProcessor()
	})
}

func NewMoveProcessor() *MoveProcessor {
	return &MoveProcessor{
		config: &MoveConfig{},
	}
}

func (d *MoveProcessor) Config() interface{} {
	return d.config
}

func (d *MoveProcessor) Process(jsonObj *simplejson.Json) {
	if d.config == nil {
		return
	}
	for _, t := range d.config.Target {
		from := t.From
		to := t.To

		upperPath, lastQuery := util.GetQueryUpperPaths(from)
		upperObj := jsonObj.GetPath(upperPath...)

		tmp := upperObj.Get(lastQuery)
		toPaths := util.GetQueryPaths(to)
		jsonObj.SetPath(toPaths, tmp)

		upperObj.Del(lastQuery)
	}
}
