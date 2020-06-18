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
	cmdList   bool
	listAll   bool
	cmdUpdate bool
	cmdRemove bool
)

func main() {
	log.Debug("start")

	parser := flags.New("ucsh-config")
	parser.BoolVar(&cmdList, "l", false,
		"list settings (exclude default values unless -a)")
	parser.BoolVar(&listAll, "a", false,
		"list all settings (only useful with -l)")
	parser.BoolVar(&cmdUpdate, "u", false,
		"update setting: -u section.key val")
	parser.BoolVar(&cmdRemove, "r", false,
		"remove setting: -r section.key")
	flags.Parse(parser, args)

	sh := ucsh.New()
	cmd.Configure(sh)
	cmd.UserConfig(sh)

	if cmdList {
		if cmdUpdate || cmdRemove {
			flags.ShowHelp(parser,
				"-l, -u and -r are mutually exclusive")
		}
		filter := parser.Arg(0)
		list(sh, filter)
	} else if cmdUpdate {
		if cmdList || cmdRemove {
			flags.ShowHelp(parser,
				"-l, -u and -r are mutually exclusive")
		}
		update(sh, parser.Arg(0), parser.Arg(1))
	} else if cmdRemove {
		if cmdList || cmdUpdate {
			flags.ShowHelp(parser,
				"-l, -u and -r are mutually exclusive")
		}
	} else {
		flags.ShowHelp(parser,
			"no action, try with -l (list), -u (update) or -r (remove)")
		os.Exit(2)
	}

	log.Debug("end")
}
