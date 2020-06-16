// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cmd

import (
	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/ucsh"
)

func Main(sh *ucsh.UCSh) {
	log.Debugf("main %s", sh)
	sh.Check()
}
