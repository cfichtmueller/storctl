// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"errors"
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

var (
	configDir  string
	configFile string
)

type Config struct {
	CurrentContext string     `yaml:"current-context"`
	current        *Context   `yaml:"-"`
	Contexts       []*Context `yaml:"contexts"`
}

func newConfig() *Config {
	return &Config{
		Contexts: make([]*Context, 0),
	}
}

func (c *Config) CreateContext(name, server, apiKey string) error {
	if c.contextExists(name) {
		return fmt.Errorf("context already exists")
	}
	c.Contexts = append(c.Contexts, &Context{
		Name:   name,
		Server: server,
		ApiKey: apiKey,
	})
	return nil
}

func (c *Config) DeleteContext(name string) error {
	if !c.contextExists(name) {
		return fmt.Errorf("context doesn't exist")
	}
	if c.CurrentContext == name {
		return fmt.Errorf("cannot delete current context")
	}
	contexts := make([]*Context, 0)
	for _, ctx := range c.Contexts {
		if ctx.Name != name {
			contexts = append(contexts, ctx)
		}
	}
	c.Contexts = contexts
	return nil
}

func (c *Config) SetCurrentContext(name string) error {
	for _, ctx := range c.Contexts {
		if ctx.Name == name {
			c.CurrentContext = ctx.Name
			c.current = ctx
			return nil
		}
	}
	return fmt.Errorf("context doesn't exist")
}

func (c *Config) GetCurrentContext() *Context {
	return c.current
}

func (c *Config) contextExists(name string) bool {
	for _, ctx := range c.Contexts {
		if ctx.Name == name {
			return true
		}
	}
	return false
}

type Context struct {
	Name   string `yaml:"name"`
	Server string `yaml:"server"`
	ApiKey string `yaml:"api-key"`
}

func Load() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot determine user home dir: %v", err)
	}
	configDir = path.Join(home, ".storctl")
	configFile = path.Join(configDir, "config")

	exists, err := configDirExists()
	if err != nil {
		return nil, err
	}
	if !exists {
		if err := os.Mkdir(configDir, 0700); err != nil {
			return nil, fmt.Errorf("cannot create config dir: %v", err)
		}
	}

	if _, err := os.Stat(configFile); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return newConfig(), nil
		}
		return nil, err
	}

	b, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %v", err)
	}

	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("invalid config file: %v", err)
	}

	if c.CurrentContext != "" {
		for _, ctx := range c.Contexts {
			if ctx.Name == c.CurrentContext {
				c.current = ctx
				break
			}
		}
	}

	return &c, nil
}

func Save(c *Config) error {
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("unable to marshal config: %v", err)
	}
	if err := os.WriteFile(path.Join(configDir, "config"), b, 0600); err != nil {
		return fmt.Errorf("unable to write config file: %v", err)
	}
	return nil
}

func configDirExists() (bool, error) {
	if _, err := os.Stat(configDir); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
