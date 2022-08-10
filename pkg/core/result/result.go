package result

import "bailuoxi66/go-loggie/pkg/core/api"

type DefaultResult struct {
	err    error
	status api.Status
	es     []api.Event
}

func Success() DefaultResult {
	return NewResult(api.SUCCESS)
}

func Fail(err error) DefaultResult {
	return DefaultResult{
		err:    err,
		status: api.FAIL,
	}
}

func NewResult(state api.Status) DefaultResult {
	return DefaultResult{
		status: state,
	}
}

func (dr DefaultResult) WithError(err error) DefaultResult {
	dr.err = err
	return dr
}

func (dr DefaultResult) Events() []api.Event {
	return dr.es
}

func (dr DefaultResult) Status() api.Status {
	return dr.status
}

func (dr DefaultResult) Batch() api.Batch {
	return nil
}

func (dr DefaultResult) ChangeStatusTo(status api.Status) {
	dr.status = status
}

func (dr DefaultResult) Error() error {
	return dr.err
}
