package interceptor

import "bailuoxi66/go-loggie/pkg/core/api"

const DefaultOrder = 999

type Extension interface {
	Order() int
	BelongTo() (componentTypes []string)
	IgnoreRetry() bool
}

type ExtensionConfig struct {
	Order       int      `yaml:"order,omitempty" default:"900"`
	BelongTo    []string `yaml:"belongTo,omitempty"`
	IgnoreRetry bool     `yaml:"ignoreRetry,omitempty" default:"true"`
}

type SortableInterceptor []api.Interceptor

func (si SortableInterceptor) Len() int {
	return len(si)
}

func (si SortableInterceptor) Less(i, j int) bool {
	i1 := si[i]
	i2 := si[j]
	var o1, o2 int
	{
		e1, ok := i1.(Extension)
		if ok {
			o1 = e1.Order()
		} else {
			o1 = DefaultOrder
		}
	}
	{
		e2, ok := i2.(Extension)
		if ok {
			o2 = e2.Order()
		} else {
			o2 = DefaultOrder
		}
	}
	return o1 < o2
}

func (si SortableInterceptor) Swap(i, j int) {
	si[i], si[j] = si[j], si[i]
}
