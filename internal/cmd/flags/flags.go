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

var help bool
var debug bool
var version bool

func New(name string) *flag.FlagSet {
	f := flag.NewFlagSet(name, flag.ExitOnError)
	f.BoolVar(&help, "h", false, "show this usage information and exit")
	f.BoolVar(&debug, "d", false, "enable debug log")
	f.BoolVar(&version, "V", false, "show version information and exit")
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
