// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package user

import (
	"github.com/jrmsdev/ucsh/internal/env"
	"github.com/jrmsdev/ucsh/internal/log"
)

type User struct {
	env *env.Env
}

func New(e *env.Env) (*User, error) {
	log.Debug("new")
	if err := e.Validate(); err != nil {
		return nil, err
	}
	return &User{e}, nil
}

func (u *User) String() string {
	return "FIXME"
}
