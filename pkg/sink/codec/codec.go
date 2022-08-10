package codec

import (
	"bailuoxi66/go-loggie/pkg/core/api"
	"bailuoxi66/go-loggie/pkg/core/log"
)

type SinkCodec interface {
	SetCodec(c Codec)
}

type Codec interface {
	Init()
	Encode(event api.Event) (*Result, error)
}

type Result struct {
	Raw []byte
	Lookup
}

type Lookup func(paths ...string) (interface{}, error)

type Factory func() Codec

var center = make(map[string]Factory)

func Register(name string, factory Factory) {
	_, ok := center[name]
	if ok {
		log.Panic("codec %s is duplicated", name)
	}

	center[name] = factory
}

func Get(name string) (Codec, bool) {
	f, ok := center[name]
	if !ok {
		return nil, ok
	}
	return f(), ok
}
