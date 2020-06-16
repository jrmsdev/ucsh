// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/ucsh/internal/cmd"
	"github.com/jrmsdev/ucsh/internal/cmd/flags"
	"github.com/jrmsdev/ucsh/internal/config"
	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/ucsh"
)

var args []string
var def *config.Config

func init() {
	args = os.Args[1:]
	def = config.New()
}

var (
	cmdList bool
	listAll bool
)

func main() {
	log.Debug("start")

	parser := flags.New("ucsh-config")
	parser.BoolVar(&cmdList, "l", false, "list settings (exclude default values unless -a)")
	parser.BoolVar(&listAll, "a", false, "list all settings (only useful with -l)")
	flags.Parse(parser, args)

	sh := ucsh.New()
	cmd.Configure(sh)
	cmd.UserConfig(sh)

	if cmdList {
		filter := parser.Arg(0)
		list(sh, filter)
	} else {
		flags.ShowHelp(parser)
	}

	log.Debug("end")
}
