package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type server struct {
	Listen string `yaml:"listen"`
	Port   int    `yaml:"port"`
}

type client struct {
	Server string `yaml:"server"`
	Port   int    `yaml:"port"`
}

type config struct {
	Type   string `yaml:"type"`
	Server server `yaml:"server"`
	Client client `yaml:"client"`
}

var Config config

func InitConfig(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, &Config)
	if err != nil {
		return err
	}
	return nil
}
