package component

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type MysqlConfig struct {
	DNS string `yaml:"dns"`
}

type TopicConfig struct {
	Name    string `yaml:"name"`
	GroupID string `yaml:"group_id"`
}

type KafkaConfig struct {
	Brokers           string      `yaml:"brokers"`
	ShoppingDataTopic TopicConfig `yaml:"shopping_data_topic"`
}

type LogConfig struct {
	Output   int    `yaml:"output"`
	Folder   string `yaml:"folder"`
	FileName string `yaml:"file_name"`
	Format   int    `yaml:"format"`
	Level    int    `yaml:"level"`
}

type Config struct {
	Mysql  MysqlConfig `yaml:"mysql"`
	Kafka  KafkaConfig `yaml:"kafka"`
	Logger LogConfig   `yaml:"logger"`
}

func LoadConfig(fileName string) (*Config, error) {
	configData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("loading config file failed: %v", err)
	}
	conf := Config{}
	err = yaml.Unmarshal(configData, &conf)
	if err != nil {
		return nil, fmt.Errorf("reading config failed: %v", err)
	}
	return &conf, nil
}
