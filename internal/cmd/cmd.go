// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cmd

import (
	"github.com/jrmsdev/ucsh"
	"github.com/jrmsdev/ucsh/internal/log"
)

func Main(sh *ucsh.UCSh) {
	log.Debugf("main %s", sh)
	sh.Check()
}
