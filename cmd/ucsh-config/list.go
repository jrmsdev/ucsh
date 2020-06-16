// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"

	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/ucsh"
)

func list(sh *ucsh.UCSh, filter string) {
	log.Debug("list cmd")
	for k, v := range sh.Config.List(filter) {
		fmt.Printf("%s=%s\n", k, v)
	}
}
