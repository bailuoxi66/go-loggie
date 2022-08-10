package pipeline

import (
	"bailuoxi66/go-loggie/pkg/core/api"
	"bailuoxi66/go-loggie/pkg/core/cfg"
	"bailuoxi66/go-loggie/pkg/core/event"
	"bailuoxi66/go-loggie/pkg/core/interceptor"
	"bailuoxi66/go-loggie/pkg/core/queue"
	"bailuoxi66/go-loggie/pkg/core/sink"
	"bailuoxi66/go-loggie/pkg/core/source"
	"log"
	"time"
)

type Config struct {
	Name             string        `yaml:"name" validate:"required"`
	CleanDataTimeout time.Duration `yaml:"cleanDataTimeout,omitempty" default:"5s"`

	Queue        *queue.Config        `yaml:"queue,omitempty" validate:"dive,required"`
	Interceptors []interceptor.Config `yaml:"interceptors,omitempty"`
	Sources      []source.Config      `yaml:"sources,omitempty" validate:"dive,required"`
	Sink         *sink.Config         `yaml:"sink,omitempty" validate:"dive,required"`
}

type ConfigRaw struct {
	Name             string        `yaml:"name" validate:"required"`
	CleanDataTimeout time.Duration `yaml:"cleanDataTimeout,omitempty" default:"5s"`

	Queue        cfg.CommonCfg   `yaml:"queue,omitempty" validate:"required"`
	Interceptors []cfg.CommonCfg `yaml:"interceptors,omitempty"`
	Sources      []cfg.CommonCfg `yaml:"sources,omitempty" validate:"required"`
	Sink         cfg.CommonCfg   `yaml:"sink,omitempty" validate:"required"`
}

func (cr *ConfigRaw) SetDefaults() {
	defaults, err := GetDefaultConfigRaw()
	if err != nil {
		log.Fatalf("get default pipeline config error: %v", err)
	}
	cr.Queue = cfg.MergeCommonCfg(cr.Queue, defaults.Queue, false)
	cr.Interceptors = cfg.MergeCommonCfgListByType(cr.Interceptors, defaults.Interceptors, false, false)
	cr.Sink = cfg.MergeCommonCfg(cr.Sink, defaults.Sink, false)
	cr.Sources = cfg.MergeCommonCfgListByType(cr.Sources, defaults.Sources, false, true)
}

func (cr *ConfigRaw) ToConfig() (*Config, error) {
	config := &Config{}

	config.Name = cr.Name
	config.CleanDataTimeout = cr.CleanDataTimeout

	queueConfig := queue.Config{}
	err := cfg.UnpackDefaultsAndValidate(cr.Queue, &queueConfig)
	if err != nil {
		return nil, err
	}

	queueConfig.Properties = cr.Queue.GetProperties()
	config.Queue = &queueConfig

	for _, in := range cr.Interceptors {
		interConfig := interceptor.Config{}
		err := cfg.UnpackDefaultsAndValidate(in, &interConfig)
		if err != nil {
			return nil, err
		}
		interConfig.Properties = in.GetProperties()
		config.Interceptors = append(config.Interceptors, interConfig)
	}

	sinkConfig := sink.Config{}
	err = cfg.UnpackDefaultsAndValidate(cr.Sink, &sinkConfig)
	if err != nil {
		return nil, err
	}
	sinkConfig.Properties = cr.Sink.GetProperties()
	config.Sink = &sinkConfig

	for _, sr := range cr.Sources {
		srcConfig := source.Config{}
		err := cfg.UnpackDefaultsAndValidate(sr, &srcConfig)
		if err != nil {
			return nil, err
		}
		srcConfig.Properties = sr.GetProperties()
		config.Sources = append(config.Sources, srcConfig)
	}

	return config, nil
}

type Info struct {
	Stop         bool // lazy stop signal
	PipelineName string
	SurviveChan  chan api.Batch
	Epoch        Epoch
	R            *RegisterCenter
	SinkCount    int
	EventPool    *event.Pool
}

func (cr *ConfigRaw) Validate() error {
	c, err := cr.ToConfig()
	if err != nil {
		return err
	}
	return c.Validate()
}

func (c *Config) Validate() error {
	return NewPipeline().validate(c)
}
