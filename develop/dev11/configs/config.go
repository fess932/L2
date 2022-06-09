package configs

type Config struct {
	Addr string
}

func NewConfig() *Config {
	return &Config{":8080"}
}
