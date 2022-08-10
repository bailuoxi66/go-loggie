package spi

import "bailuoxi66/go-loggie/pkg/core/api"

type ComponentListener interface {
	Name() string
	Stop()
}

type QueueListener interface {
	ComponentListener
	BeforeQueueConvertBatch(events []api.Event)
}
