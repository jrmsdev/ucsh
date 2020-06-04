// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"os"

	"github.com/jrmsdev/ucsh/internal/log"
)

type Env struct {
}

func New() *Env {
	log.Debug("new")
	return &Env{}
}

func (e *Env) Debug() {
	for _, e := range os.Environ() {
		log.Debugf("env %s", e)
	}
}

func (e *Env) Validate() error {
	log.Debug("validate")
	return nil
}
