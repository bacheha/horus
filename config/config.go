package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func (c *Config) SetBindings(bindings []string) {
	for _, binding := range bindings {
		c.v.BindEnv(binding)
	}
}

func (c *Config) Load(name string, configType string, path string, o interface{}) error {
	c.v.SetConfigName(name)
	c.v.SetConfigType(configType)
	c.v.AddConfigPath(path)
	if err := c.v.ReadInConfig(); err != nil {
		return err
	}
	if err := c.v.Unmarshal(&o); err != nil {
		return err
	}
	c.v.AutomaticEnv()
	return nil
}

func New(prefix string) *Config {
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return &Config{v: v}
}
