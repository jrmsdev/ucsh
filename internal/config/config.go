// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/jrmsdev/ucsh/internal/log"
)

type section interface {
	kmap() map[string]*string
	setDefaults()
}

type Config struct {
	section   map[string]section
	User      *User      `json:"user,omitempty"`
	Container *Container `json:"container,omitempty"`
}

func New() *Config {
	c := &Config{
		Container: newContainer(),
		User:      newUser(),
	}
	c.section = map[string]section{
		"container": c.Container,
		"user": c.User,
	}
	return c
}

type ucsh struct {
	D *Config `json:"ucsh,omitempty"`
}

func (c *Config) setDefaults() {
	for sn, s := range c.section {
		log.Debugf("set defaults %s section", sn)
		s.setDefaults()
	}
}

func (c *Config) Load(filename string, fh io.Reader) error {
	log.Debugf("load %s", filename)
	c.setDefaults()
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

func (c *Config) Save(filename string, fh io.Writer) error {
	log.Debugf("save %s", filename)
	obj := &ucsh{D: c}
	blob, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		log.Error(err)
		return err
	}
	if err := ioutil.WriteFile(filename, blob, 0644); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
