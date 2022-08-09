package batch

import (
	"bailuoxi66/go-loggie/pkg/core/api"
	"sync"
	"time"
)

type DefaultBatch struct {
	Es        []api.Event
	size      int
	startTime time.Time
	meta      map[string]interface{}
}

func (db *DefaultBatch) Meta() map[string]interface{} {
	return db.meta
}

func (db *DefaultBatch) Events() []api.Event {
	return db.Es
}

func (db *DefaultBatch) Release() {
	ReleaseBatch(db)
}

func (db *DefaultBatch) append(e api.Event) {
	db.Es = append(db.Es, e)
	db.size++
}

var pool = sync.Pool{
	New: func() interface{} {
		return &DefaultBatch{}
	},
}

func NewBatchWithEvents(events []api.Event) *DefaultBatch {
	b := pool.Get().(*DefaultBatch)
	*b = DefaultBatch{
		Es:        events,
		startTime: time.Now(),
		meta:      make(map[string]interface{}),
	}
	return b
}

func NewBatch() *DefaultBatch {
	b := pool.Get().(*DefaultBatch)
	*b = DefaultBatch{
		startTime: time.Now(),
		meta:      make(map[string]interface{}),
	}
	return b
}

func ReleaseBatch(b *DefaultBatch) {
	*b = DefaultBatch{} // clear batch
	pool.Put(b)
}
