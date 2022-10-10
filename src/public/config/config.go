package config

import (
	"bytes"
	cf "github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"gopkg.in/yaml.v3"
	"os"
)

type Configuration struct {
	App struct {
		Port       uint `yaml:"port"`
		Datasource struct {
			Driver   string `yaml:"driver"`
			Host     string `yaml:"host" `
			Port     uint   `yaml:"port,omitempty"`
			Username string `yaml:"username,omitempty"`
			Password string `yaml:"password,omitempty"`
			Database string `yaml:"database,omitempty"`
		}
		Redis struct {
			Host string `yaml:"host"`
			Port uint   `yaml:"port"`
		}
	}
}

func LoadConfig() *Configuration {
	var c Configuration
	cf.WithOptions(cf.ParseEnv)
	cf.AddDriver(yamlv3.Driver)

	path, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	err = cf.LoadFiles(path + "/config/default.yml")

	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)

	_, err = cf.DumpTo(buf, cf.Yaml)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(buf.Bytes(), &c); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	return &c
}
