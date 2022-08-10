package memory

import "time"

type Config struct {
	BatchSize          int           `yaml:"batchSize" default:"2048"`
	BatchBufferFactor  int           `yaml:"batchBufferFactor" default:"2"` // BufferSize = BatchSize * BatchBufferFactor
	BatchBytes         int64         `yaml:"batchBytes" default:"33554432"` // default:32MB
	BatchAggMaxTimeout time.Duration `yaml:"batchAggTimeout" default:"1s"`
	CleanDataTimeout   time.Duration `yaml:"cleanDataTimeout" default:"5s"`
}
