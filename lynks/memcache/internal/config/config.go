package config

import (
	"github.com/spf13/viper"
)

const (
	confType = "yaml"
	confName = "config-memcache"
	confPath = "./"
)

type Config struct {
	App   App   `yaml:"app"`
	Api   Api   `yaml:"api"`
	Log   Log   `yaml:"log"`
	Redis Redis `yaml:"redis"`
}
type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}
type Api struct {
	Port int
	Host string
}
type Log struct {
	Level string `yaml:"level"`
}
type Redis struct {
	Host              string `yaml:"host"`
	Port              string `yaml:"port"`
	Password          string `yaml:"password"`
	ValueLifeTimeHour int    `yaml:"valueLifeTimeHour"`
}

func Init() (*Config, error) {
	viper.SetConfigType(confType)
	viper.SetConfigName(confName)
	viper.AddConfigPath(confPath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := Config{}
	err = viper.UnmarshalKey("configuration", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
