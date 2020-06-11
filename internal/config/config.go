// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/jrmsdev/ucsh/internal/log"
)

type Config struct {
	obj  *ucsh
	User *User `json:"user,omitempty"`
}

type ucsh struct {
	D *Config `json:"ucsh,omitempty"`
}

func New() *Config {
	c := new(Config)
	c.User = newUser()
	c.obj = &ucsh{D: c}
	return c
}

func (c *Config) Load(name string, fh io.Reader) error {
	log.Debugf("load %s", name)
	blob, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Error(err)
		return err
	}
	if err := json.Unmarshal(blob, c.obj); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
