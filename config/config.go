package config

import (
	"github.com/galehuang/demo-project/utils/yaml_util"
)

type Config struct {
	DB DBConfig `yaml:"db"`

}

type DBConfig struct {
	Main MainDBConfig `yaml:"main"`
}

type MainDBConfig struct {
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Addr string `yaml:"addr"`
	Name string `yaml:"name"`
	Charset string `yaml:"charset"`
}

var gConfig *Config

func init(){
	gConfig = &Config{}
}

func GetConfig() *Config{
	return gConfig
}

func (c *Config) LoadConfigFromProfile(profilePath string) error{
	return yamlutil.UnmarshalFromFile(profilePath, c, true)
}


