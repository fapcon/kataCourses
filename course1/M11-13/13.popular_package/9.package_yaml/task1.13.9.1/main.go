package main

import (
	"fmt"
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

	z := []Config{
		{
			Server: Server{Port: "8080"},
			Db:     Db{Host: "localhost", Port: "5432", User: "admin", Password: "password123"},
		},
	}

	/*
		_, err := getYAML(r)
		if err != nil {
			panic(err)
		}

	*/
	res, _ := getYAML(z)
	_, err := fmt.Println(res)
	if err != nil {
		panic(err)
	}

}

func getYAML(config []Config) (string, error) {
	b, err := yaml.Marshal(config)
	if err != nil {
		panic(err)
	}
	return string(b), err
}
