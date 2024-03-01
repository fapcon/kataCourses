package main

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server Server `yaml:"geogrpc"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func main() {

}

func getConfigFromYAML(data []byte) (Config, error) {
	var config Config
	var ni Config
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return ni, err
	}
	return config, nil
}
