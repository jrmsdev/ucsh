// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/ucsh/internal/cmd"
	"github.com/jrmsdev/ucsh/internal/cmd/flags"
	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/ucsh"
)

func main() {
	log.Debug("start")

	parser := flags.New("ucsh")
	flags.Parse(parser)

	sh := ucsh.New()
	cmd.Configure(sh)
	cmd.UserLoad(sh)
	cmd.UserConfig(sh)

	log.Debug("end")
}
