// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

type Env struct {
}

func New() *Env {
	return &Env{}
}

func (e *Env) Validate() error {
	return nil
}
