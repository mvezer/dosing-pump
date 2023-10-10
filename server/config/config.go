package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

const (
	CONFIG_FILE = "config.yaml"
)

type Config struct {
	MlPumpInterval int `yaml:"timeOut"`
	PumpSpeed      int `yaml:"pumpSpeed"`
	Schedule       []struct {
		PumpId   int    `yaml:"pumpId"`
		AmountUl int    `yaml:"amount"`
		Cron     string `yaml:"cron"`
	}
}

func LoadConfig(fileName string) (*Config, error) {
	configYaml, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return ParseConfig(configYaml)
}

func SaveConfig(config *Config) error {
	configYaml, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(CONFIG_FILE, configYaml, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ParseConfig(yamlData []byte) (*Config, error) {
	c := new(Config)
	err := yaml.Unmarshal(yamlData, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
