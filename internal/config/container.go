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
	return &Container{
		Engine: "schroot",
	}
}

func (c *Container) list() map[string]string {
	return map[string]string{
		"engine": c.Engine,
	}
}
