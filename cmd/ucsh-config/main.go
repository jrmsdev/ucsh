// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"

	"github.com/jrmsdev/ucsh/internal/cmd"
	"github.com/jrmsdev/ucsh/internal/cmd/flags"
	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/ucsh"
)

func main() {
	log.Debug("start")

	parser := flags.New("ucsh-config")
	flags.Parse(parser)

	sh := ucsh.New()
	cmd.Configure(sh)
	cmd.UserConfig(sh)

	for k, v := range sh.Config.List("") {
		fmt.Printf("%s=%s\n", k, v)
	}

	log.Debug("end")
}
