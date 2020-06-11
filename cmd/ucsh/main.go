// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"
	"path/filepath"

	"github.com/jrmsdev/ucsh"
	"github.com/jrmsdev/ucsh/internal/cmd"
	"github.com/jrmsdev/ucsh/internal/log"
)

func userCfgFile() string {
	d, err := os.UserConfigDir()
	if err != nil {
		return ""
	}
	return filepath.Join(d, "ucsh.cfg")
}

var cfgfiles = []string{
	filepath.FromSlash("/etc/ucsh.cfg"),
	filepath.FromSlash("/usr/local/etc/ucsh.cfg"),
	userCfgFile(),
}

func main() {
	log.Debug("start")
	sh := ucsh.New()
	log.Debug(sh)
	cfgerr := false
	for _, fn := range cfgfiles {
		if fn == "" {
			continue
		}
		fh, err := os.Open(fn)
		if err != nil {
			log.Error(err)
			if !os.IsNotExist(err) {
				cfgerr = true
			}
		} else {
			if err := sh.Config.Load(fh); err != nil {
				log.Error(err)
				cfgerr = true
			}
		}
	}
	if cfgerr {
		sh.Fail("config error")
	}
	cmd.Main(sh)
	log.Debug("end")
}
