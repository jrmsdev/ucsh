// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"os"

	"github.com/jrmsdev/ucsh/internal/log"
)

var Debug bool

func New(name string) *flag.FlagSet {
	f := flag.NewFlagSet(name, flag.ExitOnError)
	f.BoolVar(&Debug, "d", false, "enable debug log")
	return f
}

var osArgs []string

func init() {
	osArgs = os.Args[1:]
}

func Parse(f *flag.FlagSet) {
	f.Parse(osArgs)
	if Debug {
		log.SetDebug(true)
	}
}
