package sink

import (
	"bailuoxi66/go-loggie/pkg/core/api"
	"bailuoxi66/go-loggie/pkg/core/interceptor"
	"fmt"
	"sort"
)

const (
	sinkInterceptorCode = "abstract-sink-interceptor"
)

type Interceptor interface {
	api.Interceptor
	Intercept(invoker Invoker, invocation Invocation) api.Result
}

type AbstractInterceptor struct {
	DoName      func() string
	DoIntercept func(invoker Invoker, invocation Invocation) api.Result
}

func (ai *AbstractInterceptor) Intercept(invoker Invoker, invocation Invocation) api.Result {
	return ai.DoIntercept(invoker, invocation)
}

func (ai *AbstractInterceptor) Init(context api.Context) {
	// ignore
}

func (ai *AbstractInterceptor) Start() {
	// ignore
}

func (ai *AbstractInterceptor) Stop() {
	// ignore
}

func (ai *AbstractInterceptor) Category() api.Category {
	return api.INTERCEPTOR
}

func (ai *AbstractInterceptor) Type() api.Type {
	return api.Type(ai.DoName())
}

func (ai *AbstractInterceptor) String() string {
	return fmt.Sprintf("%s/%s", ai.Category(), ai.Type())
}

func (ai *AbstractInterceptor) Config() interface{} {
	return nil
}

type SortableInterceptor []Interceptor

func (si SortableInterceptor) Len() int {
	return len(si)
}

func (si SortableInterceptor) Less(i, j int) bool {
	i1 := si[i]
	i2 := si[j]
	var o1, o2 int
	{
		e1, ok := i1.(interceptor.Extension)
		if ok {
			o1 = e1.Order()
		} else {
			o1 = interceptor.DefaultOrder
		}
	}
	{
		e2, ok := i2.(interceptor.Extension)
		if ok {
			o2 = e2.Order()
		} else {
			o2 = interceptor.DefaultOrder
		}
	}
	return o1 < o2
}

func (si SortableInterceptor) Swap(i, j int) {
	si[i], si[j] = si[j], si[i]
}

func (si SortableInterceptor) Sort() {
	sort.Sort(si)
}
