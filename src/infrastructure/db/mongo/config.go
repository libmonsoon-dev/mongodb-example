package mongo

type Config struct {
	Host         string
	Port         int
	DataBaseName string
}

func DefaultConfig() *Config {
	return &Config{
		"localhost",
		27017,
		"mongo-example",
	}
}

func TestConfig() *Config {
	return &Config{
		"localhost",
		27027,
		"tests",
	}
}
