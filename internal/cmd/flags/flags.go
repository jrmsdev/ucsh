// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"fmt"
	"os"

	"github.com/jrmsdev/ucsh"
	"github.com/jrmsdev/ucsh/internal/log"
)

var (
	help bool = false
	debug bool = false
	version bool = false
	ConfigFile string = "config.json"
)

func New(name string) *flag.FlagSet {
	f := flag.NewFlagSet(name, flag.ExitOnError)
	f.BoolVar(&help, "h", help, "show this usage information and exit")
	f.BoolVar(&debug, "d", debug, "enable debug log")
	f.BoolVar(&version, "V", version, "show version information and exit")
	f.StringVar(&ConfigFile, "c", ConfigFile, "config file `name`")
	return f
}

func Parse(f *flag.FlagSet, args []string) {
	f.Parse(args)
	if debug {
		log.SetDebug(true)
	}
	if help {
		ShowHelp(f)
	} else if version {
		fmt.Fprintf(f.Output(), "%s version %s\n", f.Name(), ucsh.Version())
		os.Exit(2)
	}
}

func ShowHelp(f *flag.FlagSet) {
	fmt.Fprintf(f.Output(), "Usage for %s:\n", f.Name())
	f.PrintDefaults()
	os.Exit(2)
}
