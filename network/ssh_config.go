package network

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// Config is a network configuration
type Config struct {
	Hosts []HostConfig
}

// HostConfig is a host configuration
type HostConfig struct {
	User, IP string
	Nodes    []string
}

// InitConfig returns a network configuration from `cfgpath`
func InitConfig(cfgpath string) (*Config, error) {
	bytes, err := ioutil.ReadFile(cfgpath)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}