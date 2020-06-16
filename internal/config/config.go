// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/jrmsdev/ucsh/internal/log"
)

type Config struct {
	User *User `json:"user,omitempty"`
	Container *Container `json:"container,omitempty"`
}

func New() *Config {
	return &Config{
		User: newUser(),
		Container: newContainer(),
	}
}

type ucsh struct {
	D *Config `json:"ucsh,omitempty"`
}

func (c *Config) Load(name string, fh io.Reader) error {
	log.Debugf("load %s", name)
	blob, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Error(err)
		return err
	}
	obj := &ucsh{D: c}
	if err := json.Unmarshal(blob, obj); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (c *Config) List(prefix string) map[string]string {
	log.Debugf("list '%s'", prefix)
	l := make(map[string]string)
	if prefix == "" {
		for k, v := range c.ls(c.User.list(), "") {
			l["user." + k] = v
		}
	}
	return l
}

func (c *Config) ls(src map[string]string, filter string) map[string]string {
	if filter != "" {
		filter = strings.TrimSpace(filter) + "."
	}
	l := make(map[string]string)
	for k, v := range src {
		if filter == "" || strings.HasPrefix(k, filter) {
			l[k] = v
		}
	}
	return l
}
