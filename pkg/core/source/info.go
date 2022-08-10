package source

import "bailuoxi66/go-loggie/pkg/core/api"

type Info struct {
	Source       api.Source
	Queue        api.Queue
	Interceptors []Interceptor
}
