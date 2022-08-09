package context

import (
	"bailuoxi66/go-loggie/pkg/core/api"
	"bailuoxi66/go-loggie/pkg/core/cfg"
)

func NewContext(name string, typename api.Type, category api.Category, properties cfg.CommonCfg) api.Context {
	return &DefaultContext{
		name:       name,
		category:   category,
		typename:   typename,
		properties: properties,
	}
}

type DefaultContext struct {
	name       string
	category   api.Category
	typename   api.Type
	properties cfg.CommonCfg
}

func (c *DefaultContext) Properties() cfg.CommonCfg {
	return c.properties
}

func (c *DefaultContext) Category() api.Category {
	return c.category
}

func (c *DefaultContext) Type() api.Type {
	return c.typename
}

func (c *DefaultContext) Name() string {
	return c.name
}
