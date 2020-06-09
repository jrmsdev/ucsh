// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg

import (
	"github.com/jrmsdev/ucsh/internal/log"
)

// user

type userCfg struct {
	Shell string `json:"shell,omitempty"`
}

func (u *userCfg) debug() {
	log.Debugf("user.shell: %s", u.Shell)
}

var User = &userCfg{
	Shell: "/bin/sh",
}

// json schema

type config struct {
	User *userCfg `json:"user,omitempty"`
}

type ucsh struct {
	D *config `json:"ucsh,omitempty"`
}

var c = &ucsh{D: &config{
	User: User,
}}

// debug

func debug() {
	User.debug()
}
