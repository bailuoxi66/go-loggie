package pipeline

import (
	"strconv"
	"strings"
	"time"
)

type Epoch struct {
	PipelineName string
	ReloadCount  int
	StartTime    time.Time
}

func NewEpoch(pipelineName string) Epoch {
	return Epoch{
		PipelineName: pipelineName,
		ReloadCount:  0,
		StartTime:    time.Now(),
	}
}

func (e Epoch) IsEmpty() bool {
	return e.StartTime.IsZero()
}

func (e Epoch) Increase() {
	e.ReloadCount++
	e.StartTime = time.Now()
}

func (e Epoch) Equal(ae Epoch) bool {
	if e.PipelineName != ae.PipelineName {
		return false
	}
	if e.ReloadCount != ae.ReloadCount {
		return false
	}
	return e.StartTime.Equal(ae.StartTime)
}

func (e Epoch) String() string {
	var es strings.Builder
	es.Grow(64)
	es.WriteString(e.PipelineName)
	es.WriteString(":")
	es.WriteString(strconv.Itoa(e.ReloadCount))
	es.WriteString(":")
	es.WriteString(e.StartTime.Format("2006-01-02 15:04:05"))
	return es.String()
}
