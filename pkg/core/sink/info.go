package sink

import "bailuoxi66/go-loggie/pkg/core/api"

type Info struct {
	Sink         api.Sink
	Queue        api.Queue
	Interceptors []Interceptor
}
