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
	cfg := sh.Config.List(filter)
	for k, v := range def.List(filter) {
		if listAll || cfg[k] != v {
			fmt.Printf("%s=%s\n", k, cfg[k])
		}
	}
}
