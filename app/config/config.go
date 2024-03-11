package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"db"`
	Server struct {
		AppPort string `yaml:"port"`
	} `yaml:"server"`
}

var Cfg *Config

func init() {
	var err error
	Cfg, err = ReadConfigFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
}
func ReadConfigFile(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
