package invoke

import (
	"bailuoxi66/go-loggie/pkg/core/api"
	"bailuoxi66/go-loggie/pkg/core/event"
)

type DefaultInvocation struct {
	producer  api.Producer
	consumers []api.Consumer
	selector  api.Selector
	event     *event.DefaultEvent
}

func NewDefaultInvocation(producer api.Producer, consumers []api.Consumer, selector api.Selector) *DefaultInvocation {
	return &DefaultInvocation{
		producer:  producer,
		consumers: consumers,
		selector:  selector,
	}
}

func (di *DefaultInvocation) Producer() api.Producer {
	return di.producer
}

func (di *DefaultInvocation) Consumers() []api.Consumer {
	return di.consumers
}

func (di *DefaultInvocation) Selector() api.Selector {
	return di.selector
}

func (di *DefaultInvocation) Event() api.Event {
	return di.event
}

func (di *DefaultInvocation) AppendConsumer(consumer api.Consumer) {
	if di.consumers == nil {
		di.consumers = make([]api.Consumer, 0)
	}
	di.consumers = append(di.consumers, consumer)
}
