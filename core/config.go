package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type etcd struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Auth bool `yaml:"auth"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}


type config struct {
	Etcd etcd `yaml:"etcd"`
}

var Config *config

// 初始化配置文件
func NewConfig(path string) (*config, error) {
	var err error
	var content []byte
	Config = new(config)
	content, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, Config)
	if err != nil {
		return nil, err
	}
	return Config, nil
}