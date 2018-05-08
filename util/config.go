package util

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Host    string  `yaml:"host"`
	Port    string  `yaml:"port"`
	Mongodb Mongodb `yaml:"mongodb"`
}

type Mongodb struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DB   string `yaml:"db"`
}

var config *Config

func InitConf() *Config {
	configFile, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config
}

func GetConf() *Config {
	return config
}
