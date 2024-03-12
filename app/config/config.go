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
	Jwt struct {
		ApiSecret     string `yaml:"api_secret"`
		TokenLifeSpan string `yaml:"token_hour_span"`
	}
}

var Cfg *Config

// 用来获取yaml配置文件信息并且放在Cfg变量中，在程序加载config包的时候进行初始化
func init() {
	var err error
	Cfg, err = ReadConfigFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
}

// ReadConfigFile 用来读取yaml配置文件中的信息
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
