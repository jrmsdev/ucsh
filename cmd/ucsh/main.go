// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/ucsh/internal/cmd"
	"github.com/jrmsdev/ucsh/internal/cmd/flags"
	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/ucsh"
)

var args []string

func init() {
	args = os.Args[1:]
}

func main() {
	log.Debug("start")

	parser := flags.New("ucsh")
	flags.Parse(parser, args)

	sh := ucsh.New()
	cmd.Configure(sh)
	cmd.UserLoad(sh)
	cmd.UserConfig(sh)

	log.Debug("end")
}
