// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ucsh

import (
	"os"

	"github.com/jrmsdev/ucsh/internal/cfg"
	"github.com/jrmsdev/ucsh/internal/env"
	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/user"
)

func Main() {
	log.Debug("start")
	if err := cfg.Init(); err != nil {
		log.Panic(err)
	}
	cfg.Debug()
	e := env.New()
	u, err := user.New(e)
	if err != nil {
		log.Panic(err)
	}
	log.Print(u)
	for _, e := range os.Environ() {
		log.Print(e)
	}
	log.Debug("end")
}
