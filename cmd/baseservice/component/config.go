package component

import (
	"fmt"
	"os"

	"github.com/Jungle20m/golang-base/general"
	"gopkg.in/yaml.v2"
)

const (
	FILE_CONFIG_LOCATION = "cmd\\baseservice\\config.yaml"
)

type AppConfig struct {
	Host string `yaml:"host"`
	Port int32  `yaml:"port"`
}

type Config struct {
	Mysql general.MysqlConfig `yaml:"mysql"`
	App   AppConfig           `yaml:"app"`
}

func LoadConfig() (*Config, error) {
	config := Config{}
	configData, err := os.ReadFile(FILE_CONFIG_LOCATION)
	if err != nil {
		return nil, fmt.Errorf("reading config failed: %v", err)
	}
	if err = yaml.Unmarshal(configData, &config); err != nil {
		return nil, fmt.Errorf("reading config failed: %v", err)
	}
	return &config, nil
}
