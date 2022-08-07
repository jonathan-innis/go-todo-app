package database

import (
	"fmt"

	"github.com/jonathan-innis/go-todo-app/pkg/utils"
)

type ConfigType string

const (
	ConfigTypeDefault     ConfigType = "default"
	ConfigTypeEnvironment ConfigType = "environment"
)

type Config struct {
	connectionTimeout int
	queryTimeout      int
	username          string
	password          string
	endpoint          string
	name              string
}

func NewConfig(connectionTimeout, queryTimeout int, username, password, endpoint, name string) *Config {
	return &Config{
		connectionTimeout: connectionTimeout,
		queryTimeout:      queryTimeout,
		username:          username,
		password:          password,
		endpoint:          endpoint,
		name:              name,
	}
}

func NewConfigFromEnvironment() *Config {
	return NewConfig(
		utils.GetIntEnv("CONNECTION_TIMEOUT", 5),
		utils.GetIntEnv("QUERY_TIMEOUT", 30),
		utils.GetEnv("USERNAME", ""),
		utils.GetEnv("PASSWORD", ""),
		utils.GetEnv("ENDPOINT", ""),
		utils.GetEnv("NAME", ""),
	)
}

func DefaultConfig() *Config {
	return NewConfig(
		5,
		30,
		"",
		"",
		"localhost:27017",
		"todo-app",
	)
}

func (c *Config) ConnectionString() string {
	if c.username == "" && c.password == "" {
		return fmt.Sprintf("mongodb://%s/%s", c.endpoint, c.name)
	}
	return fmt.Sprintf("mongodb://%s:%s@%s/%s", c.username, c.password, c.endpoint, c.name)
}
