// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	//~ "fmt"

	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/ucsh"
)

func update(sh *ucsh.UCSh, key, val string) {
	log.Debugf("update key '%s'", key)
	//~ cfg := sh.Config.List(filter)
	//~ for k, v := range def.List(filter) {
		//~ if listAll || cfg[k] != v {
			//~ fmt.Printf("%s=%s\n", k, cfg[k])
		//~ }
	//~ }
	sh.Config.Save("/tmp/config.json", nil)
}
