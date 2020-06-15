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
	User *User `json:"user,omitempty"`
}

func New() *Config {
	return &Config{
		User: newUser(),
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
