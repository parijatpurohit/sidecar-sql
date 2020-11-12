package config

type SQL struct {
	Dialect       string `yaml:"dialect"`
	ConnectionStr string `yaml:"connStr"`
	MaxIdle       int    `yaml:"maxIdle"`
	MaxOpen       int    `yaml:"maxOpen"`
}
type SQLConfig struct {
	SQL SQL `yaml:"sql"`
}
