package config

import (
	_ "embed"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host             string `yaml:"Host"`
	Port             int    `yaml:"Port"`
	EtcdHost         string `yaml:"EtcdHost"`
	OrderServiceName string `yaml:"OrderServiceName"`
	LogDir           string `yaml:"LogDir"`
}

var (
	instance *Config
	once     sync.Once
)

//go:embed config.yml
var data []byte

func GetConfig() *Config {
	once.Do(func() {
		instance = LoadConfig()
	})
	return instance
}

func LoadConfig() *Config {
	var conf Config

	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}

	etcdHostEnv := os.Getenv("ETCD_HOST")
	if etcdHostEnv != "" {
		conf.EtcdHost = etcdHostEnv
	}

	return &conf
}
