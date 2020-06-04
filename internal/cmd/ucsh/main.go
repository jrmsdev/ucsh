// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ucsh

import (
	"github.com/jrmsdev/ucsh/internal/cfg"
	"github.com/jrmsdev/ucsh/internal/env"
	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/user"
)

func Main() {
	log.Debug("start")
	// config
	if err := cfg.Init(); err != nil {
		log.Panic(err)
	}
	// env
	e := env.New()
	e.Debug()
	// user
	u, err := user.New(e)
	if err != nil {
		log.Panic(err)
	}
	log.Debug(u)
	log.Debug("end")
}
