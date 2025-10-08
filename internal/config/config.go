package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTP       ServerConfig     `yaml:"http"`
	GRPC       ServerConfig     `yaml:"grpc"`
	Experian   ExperianConfig   `yaml:"experian"`
	Equifax    EquifaxConfig    `yaml:"equifax"`
	TransUnion TransUnionConfig `yaml:"transunion"`
}

type ExperianConfig struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

type EquifaxConfig struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

type TransUnionConfig struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func NewConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", filename, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	// Validate HTTP config
	if err := validateServerConfig("http", cfg.HTTP); err != nil {
		return nil, err
	}

	// Validate GRPC config
	if err := validateServerConfig("grpc", cfg.GRPC); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateServerConfig(name string, sc ServerConfig) error {
	if sc.Host == "" {
		return fmt.Errorf("%s host cannot be empty", name)
	}
	if sc.Port <= 0 || sc.Port > 65535 {
		return fmt.Errorf("%s port must be between 1 and 65535", name)
	}
	return nil
}
