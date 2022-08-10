package transformer

import (
	"bailuoxi66/go-loggie/pkg/util"
	"github.com/bitly/go-simplejson"
)

type UnderRootProcessor struct {
	config *UnderRootConfig
}

type UnderRootConfig struct {
	Target []string `yaml:"target,omitempty"`
}

func init() {
	register("underRoot", func() Processor {
		return NewUnderRootProcessor()
	})
}

func NewUnderRootProcessor() *UnderRootProcessor {
	return &UnderRootProcessor{
		config: &UnderRootConfig{},
	}
}

func (d *UnderRootProcessor) Config() interface{} {
	return d.config
}

func (d *UnderRootProcessor) Process(jsonObj *simplejson.Json) {
	if d.config == nil {
		return
	}
	for _, t := range d.config.Target {
		upperPaths, key := util.GetQueryUpperPaths(t)
		upperVal := jsonObj.GetPath(upperPaths...)
		val := upperVal.Get(key)
		if valMap, err := val.Map(); err == nil {
			for k, v := range valMap {
				jsonObj.Set(k, v)
			}
		} else {
			jsonObj.Set(key, val)
		}
		upperVal.Del(key)
	}
}
