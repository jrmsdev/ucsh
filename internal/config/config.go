// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jrmsdev/ucsh/internal/log"
)

var cfgfiles = []string{
	filepath.FromSlash("/etc/ucsh.cfg"),
	filepath.FromSlash("/usr/local/etc/ucsh.cfg"),
	userCfgFile(),
}

type Config struct {
	dat  *ucsh
	User *User `json:"user,omitempty"`
}

type ucsh struct {
	D *Config `json:"ucsh,omitempty"`
}

func New() (*Config, error) {
	log.Debug("new")
	c := new(Config)
	c.User = newUser()
	c.dat = &ucsh{D: c}
	for _, fn := range cfgfiles {
		if err := c.load(fn); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Config) load(fn string) error {
	blob, err := ioutil.ReadFile(fn)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debug(err)
		} else {
			log.Error(err)
		}
	} else {
		err := json.Unmarshal(blob, c.dat)
		if err != nil {
			log.Error(err)
			return err
		}
		log.Debugf("loaded %s", fn)
	}
	return nil
}
