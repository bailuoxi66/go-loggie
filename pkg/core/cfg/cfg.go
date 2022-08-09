package cfg

import (
	"bailuoxi66/go-loggie/pkg/core/log"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type CommonCfg map[string]interface{}

type ComponentBaseConfig struct {
	Name       string    `yaml:"name"`
	Type       string    `yaml:"type" validate:"required"`
	Properties CommonCfg `yaml:",inline"`
}

// TODO
func (c CommonCfg) GetProperties() CommonCfg {
	return c
}

type Validator interface {
	Validate() error
}

func NewCommonCfg() CommonCfg {
	return make(map[string]interface{})
}

func (c CommonCfg) Put(key string, val interface{}) {
	c[key] = val
}

func (c CommonCfg) Remove(key string) {
	delete(c, key)
}

func (c CommonCfg) Get(key string) interface{} {
	return c[key]
}

// Enabled
// return false: get key 'enabled' is null or 'false'
// return true: get key 'enabled' is 'true'
func (c CommonCfg) Enabled() bool {
	if c["enabled"] == "true" {
		return true
	}
	return false
}

func (c CommonCfg) GetType() string {
	typeName, ok := c["type"]
	if !ok {
		return ""
	}
	return typeName.(string)
}

func MergeCommonCfg(base CommonCfg, from CommonCfg, override bool) CommonCfg {
	if base == nil {
		return from
	}
	if from == nil {
		return base
	}

	for k, v := range from {
		_, ok := base[k]
		if ok && !override {
			continue
		}

		base[k] = v
	}
	return base
}

// MergeCommonCfgListByType merge commonCfg list
// ignoreFromType: set ignoreFromType=true, if fromCommonCfg had a type A which does not exist in baseCommonCfg, then type A would not merged to baseCommonCfg
func MergeCommonCfgListByType(base []CommonCfg, from []CommonCfg, override bool, ignoreFromType bool) []CommonCfg {
	if len(base) == 0 {
		return from
	}
	if len(from) == 0 {
		return base
	}

	fromMap := make(map[string]CommonCfg)
	for _, v := range from {
		fromMap[v.GetType()] = v
	}

	for _, baseCfg := range base {
		typeName := baseCfg.GetType()
		fromCfg, ok := fromMap[typeName]
		if ok {
			baseCfg = MergeCommonCfg(baseCfg, fromCfg, override)
			delete(fromMap, typeName)
			continue
		}
	}

	if !ignoreFromType {
		for _, v := range fromMap {
			base = append(base, v)
		}
	}
	return base
}

func UnpackFromFileDefaultsAndValidate(path string, config interface{}) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Warn("read config error. err: %v", err)
		return err
	}

	return UnpackRawDefaultsAndValidate(content, config)
}

func UnpackFromFileDefaults(path string, config interface{}) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Warn("read config error. err: %v", err)
		return err
	}

	return UnpackRawAndDefaults(content, config)
}

func UnpackFromFile(path string, config interface{}) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Warn("read config error. err: %v", err)
		return err
	}

	return UnpackRaw(content, config)
}

func UnpackRawDefaultsAndValidate(content []byte, config interface{}) error {
	if config == nil {
		return nil
	}
	err := yaml.Unmarshal(content, config)
	if err != nil {
		return err
	}

	if err := setDefault(config); err != nil {
		return err
	}

	if err := validate(config); err != nil {
		return err
	}

	return nil
}

func UnpackRaw(content []byte, config interface{}) error {
	if config == nil {
		return nil
	}

	err := yaml.Unmarshal(content, config)
	if err != nil {
		return err
	}

	return nil
}

func UnpackRawAndDefaults(content []byte, config interface{}) error {
	if config == nil {
		return nil
	}

	err := yaml.Unmarshal(content, config)
	if err != nil {
		return err
	}

	if err := setDefault(config); err != nil {
		return err
	}
	return nil
}

func UnpackDefaultsAndValidate(properties CommonCfg, config interface{}) error {
	if properties == nil {
		return nil
	}

	out, err := yaml.Marshal(properties)
	if err != nil {
		return err
	}

	return UnpackRawDefaultsAndValidate(out, config)
}

func UnpackAndDefaults(properties CommonCfg, config interface{}) error {
	if properties == nil {
		return nil
	}

	out, err := yaml.Marshal(properties)
	if err != nil {
		return err
	}

	return UnpackRawAndDefaults(out, config)
}

func Unpack(properties CommonCfg, config interface{}) error {
	if properties == nil {
		return nil
	}

	out, err := yaml.Marshal(properties)
	if err != nil {
		return err
	}
	return UnpackRaw(out, config)
}

func Pack(config interface{}) (CommonCfg, error) {

	if config == nil {
		return nil, nil
	}

	out, err := yaml.Marshal(config)
	if err != nil {
		return nil, err
	}

	ret := make(map[string]interface{})
	err = yaml.Unmarshal(out, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func PackAndDefault(config interface{}) (CommonCfg, error) {
	err := setDefault(config)
	if err != nil {
		return nil, err
	}
	c, err := Pack(config)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func setDefault(config interface{}) error {
	if defaults.CanUpdate(config) {
		return defaults.Set(config)
	}
	return defaults.Set(config)
}

func validate(config interface{}) error {
	if config == nil {
		return nil
	}

	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		return err
	}

	if cfg, ok := config.(Validator); ok {
		return cfg.Validate()
	}
	return nil
}
