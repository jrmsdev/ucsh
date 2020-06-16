// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"fmt"
	"os"

	"github.com/jrmsdev/ucsh/internal/log"
)

var help bool
var debug bool

func New(name string) *flag.FlagSet {
	f := flag.NewFlagSet(name, flag.ExitOnError)
	f.BoolVar(&help, "h", false, "show this usage information and exit")
	f.BoolVar(&debug, "d", false, "enable debug log")
	return f
}

var osArgs []string

func init() {
	osArgs = os.Args[1:]
}

func Parse(f *flag.FlagSet) {
	f.Parse(osArgs)
	if help {
		fmt.Fprintf(f.Output(), "Usage for %s:\n", f.Name())
		f.PrintDefaults()
		os.Exit(2)
	} else if debug {
		log.SetDebug(true)
	}
}
