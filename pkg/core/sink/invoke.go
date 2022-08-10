package sink

import "bailuoxi66/go-loggie/pkg/core/api"

type Invocation struct {
	Batch api.Batch
	Sink  api.Sink
}

type Invoker interface {
	Invoke(invocation Invocation) api.Result
}

type AbstractInvoker struct {
	DoInvoke func(invocation Invocation) api.Result
}

func (ai *AbstractInvoker) Invoke(invocation Invocation) api.Result {
	return ai.DoInvoke(invocation)
}

type SubscribeInvoker struct {
}

func (si *SubscribeInvoker) Invoke(invocation Invocation) api.Result {
	return invocation.Sink.Consume(invocation.Batch)
}
