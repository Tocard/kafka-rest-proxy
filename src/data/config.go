package data

type Config struct {
	Log         Log
	KafkaBroker []string `yaml:"kafka_broker"`
	HostName    string   `yaml:"hostname"`
}

type Log struct {
	Format   string `yaml:"format"`
	LogLevel string `yaml:"log_level"`
	LogPath  string `yaml:"log_path"`
}

var YamlConfig *Config = nil

func NewConfig() *Config {
	cnf := &Config{}
	return cnf
}

func SetConfig(cfg *Config) {
	YamlConfig = cfg
}
