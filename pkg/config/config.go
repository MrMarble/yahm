package config

import (
	"io/ioutil"

	"github.com/mrmarble/ghooks/pkg/ghook"
	"gopkg.in/yaml.v2"
)

const configName = ".ghook"

type Config struct {
	Hooks map[ghook.HookType]*ghook.Hook `yaml:"hooks"`
}

// UnmarshalYAML implements the yaml.Unmarshaler interface
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Hooks = make(map[ghook.HookType]*ghook.Hook)
	type plain Config
	err := unmarshal((plain)(*c))
	if err != nil {
		return err
	}
	for hookType, hook := range c.Hooks {
		hook.Type = hookType
		c.Hooks[hookType] = hook
	}
	return nil
}

func Read(path string) (*Config, error) {
	if path == "" {
		return readFromDefault()
	}
	cfg, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.Unmarshal([]byte(cfg), &c)
	if err != nil {
		return nil, err
	}

	return &c, err
}

func readFromDefault() (*Config, error) {
	// checks if wheter the config file is .yaml or .yml
	if _, err := ioutil.ReadFile(configName + ".yaml"); err == nil {
		return Read(configName + ".yaml")
	}
	return Read(configName + ".yml")
}
