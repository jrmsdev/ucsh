// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg

import (
	"github.com/jrmsdev/ucsh/internal/log"
)

type userCfg struct {
	Shell string `json:"shell,omitempty"`
}

var User = &userCfg{
	Shell: "/bin/sh",
}

type config struct {
	User *userCfg `json:"user,omitempty"`
}

var Config = &config{
	User: User,
}

type ucsh struct {
	D *config `json:"ucsh,omitempty"`
}

var c = &ucsh{D: Config}

var cinit = false

func Init() error {
	log.Debug("init")
	if cinit {
		log.Panic("config init already done")
	}
	cinit = true
	return nil
}

func Debug() {
	log.Debugf("%#v", c)
	Init()
	log.Debugf("user.shell: %s", Config.User.Shell)
}
