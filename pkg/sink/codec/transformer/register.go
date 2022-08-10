package transformer

import (
	"bailuoxi66/go-loggie/pkg/core/log"
	"github.com/bitly/go-simplejson"
)

type Processor interface {
	Process(jsonObj *simplejson.Json)
}

type factory func() Processor

var registry = make(map[string]factory)

func register(name string, f factory) {
	_, ok := registry[name]
	if ok {
		log.Panic("transformer %s is duplicated", name)
	}
	registry[name] = f
}

func getProcessor(name string) (Processor, bool) {
	trans, ok := registry[name]
	if !ok {
		return nil, false
	}
	return trans(), true
}
