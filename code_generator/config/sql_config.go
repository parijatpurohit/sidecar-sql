package config

type SQLConfig struct {
	Dialect       string `yaml:"dialect"`
	ConnectionStr string `yaml:"connStr"`
	MaxIdle       int    `yaml:"maxIdle"`
	MaxOpen       int    `yaml:"maxOpen"`
}
