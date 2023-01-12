package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Redis struct {
		Host        string
		DeletedKeys []string `yaml:"deletedKeys"`
	}
}

func (c *Config) getConfig() *Config {
	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	if len(c.Redis.Host) == 0 {
		log.Fatalf("Redis Host should not empty")
	}

	if len(c.Redis.DeletedKeys) == 0 {
		log.Fatalf("Redis DeletedKeys should not empty")
	}
	return c
}
