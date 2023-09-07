package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	confType = "yaml"
	confName = "config-shortner"
	confPath = "./"
)

type Config struct {
	App        App        `yaml:"app"`
	Api        Api        `yaml:"api"`
	Log        Log        `yaml:"log"`
	Cache      Cache      `yaml:"cache"`
	Postgresql Postgresql `yaml:"postgresql"`
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
type Cache struct {
	Url string
}
type Postgresql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
}

func (p Postgresql) Uri() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", p.User, p.Password, p.Host, p.Port, p.DbName)
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
