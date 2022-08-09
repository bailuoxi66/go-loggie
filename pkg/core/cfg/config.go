package cfg

type SystemConfig struct {
	DbFile string `yaml:"db_file" default:"./data/loggie.db"`
}
