package source

import (
	"bailuoxi66/go-loggie/pkg/core/api"
	"bailuoxi66/go-loggie/pkg/core/result"
)

type Invocation struct {
	Event api.Event
	Queue api.Queue
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

// publish event to queue
type PublishInvoker struct {
}

func (i *PublishInvoker) Invoke(invocation Invocation) api.Result {
	invocation.Queue.In(invocation.Event)
	return result.Success()
}
