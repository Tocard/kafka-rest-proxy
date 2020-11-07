package data

type Config struct {
	Log           Log
	KafkaProducer KafkaProducer `yaml:"kafka-producer"`
}

type Log struct {
	Format   string `yaml:"format"`
	LogLevel string `yaml:"log_level"`
	LogPath  string `yaml:"log_path"`
}

type KafkaProducer struct {
	KafkaBroker []string `yaml:"kafka_broker"`
	MaxRetries  int      `yaml:"max_retries"`
}

func NewConfig() *Config {
	cnf := &Config{}
	return cnf
}
