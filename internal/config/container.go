// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/ucsh/internal/log"
)

type Container struct {
	Engine string `json:"engine,omitempty"`
}

func newContainer() *Container {
	log.Debug("new")
	return &Container{}
}

func (c *Container) setDefaults() {
	log.Debug("set defaults")
	c.Engine = "docker"
}

func (c *Container) kmap() map[string]*string {
	return map[string]*string{
		"engine": &c.Engine,
	}
}
