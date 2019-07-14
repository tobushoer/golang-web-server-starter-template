package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v3"
)

// Config yaml file format
type Config struct {
	Env   string `yaml:"env"`
	Log   Log    `yaml:"log"`
	MySQL MySQL  `yaml:"mysql"`
}

// Log yaml log key config
type Log struct {
	Pkg string `yaml:"pkg"`
}

// MySQL yaml mysql config
type MySQL struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
}

// New read config file
func New(configFile string) *Config {
	c, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panicf("Read config file %s failed: %s", configFile, err.Error())
	}
	cfg := &Config{}
	if err := yaml.Unmarshal(c, cfg); err != nil {
		log.Panicf("yaml.Unmarshal config file %s failed: %s", configFile, err.Error())
	}
	return cfg
}
