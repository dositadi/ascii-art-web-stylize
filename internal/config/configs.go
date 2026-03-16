package config

import "os"

type Config struct {
	DBUrl string
}

func (c *Config) GetDBConfig() Config {
	return Config{
		DBUrl: os.Getenv("ASCII_DB_URL"),
	}
}
