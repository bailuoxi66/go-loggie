package pipeline

import (
	"bailuoxi66/go-loggie/pkg/util"
	"github.com/pkg/errors"
)

var defaultConfigRaw ConfigRaw

func SetDefaultConfigRaw(defaults ConfigRaw) {
	defaultConfigRaw = defaults
}

func GetDefaultConfigRaw() (*ConfigRaw, error) {
	rawCopy := &ConfigRaw{}
	err := util.Clone(defaultConfigRaw, rawCopy)
	if err != nil {
		return nil, errors.WithMessage(err, "get default config failed")
	}
	return rawCopy, nil
}
