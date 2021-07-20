package conf

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"gopkg.in/yaml.v2"
)

type server struct {
	Listen     string `yaml:"listen"`
	ConfServer bool   `yaml:"conf_server"`
	Port       int    `yaml:"port"`
}

type client struct {
	Uuid string `yaml:"uuid"`
	Name string `yaml:"name"`
}

type config struct {
	Type    string   `yaml:"type"`
	Server  server   `yaml:"server"`
	Clients []client `yaml:"clients"`
}

var Config config

var ClientConfig client

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

func InitRPCConfig(path string) error {
	u, err := url.Parse(path)
	if err != nil {
		return err
	}

	query, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return err
	}

	uuid := query.Get("uuid")

	content := GetConf(u.Host, uuid)

	fmt.Println(content)

	//if err := yaml.Unmarshal(content, &ClientConfig); err != nil {
	//	return err
	//}

	return nil
}
