package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var instance *Config

// GetConfig 获取配置单例实例
func GetConfig() *Config {
	if instance != nil {
		return instance
	} else {
		c, err := loadConfig()
		if err != nil {
			fmt.Printf("config decode error, err=%v\n", err)
			panic("config decode error")
		}

		instance = c
	}

	return instance
}

func loadConfig() (*Config, error) {
	var configFile string
	flag.StringVar(&configFile, "f", "config.yaml", "config file")
	flag.Parse()

	if configFile == "" {
		fmt.Println("configFile was not provided, use default configFile ./service.yaml")
	}

	return loadYamlConfig(configFile)
}

// loadConfig 加载并解析 YAML 格式配置文件
func loadYamlConfig(configFile string) (*Config, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
