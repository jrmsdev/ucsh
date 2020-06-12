// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"
	osuser "os/user"
	"path/filepath"

	"github.com/jrmsdev/ucsh"
	"github.com/jrmsdev/ucsh/internal/cmd"
	"github.com/jrmsdev/ucsh/internal/log"
)

func userCfgFile() string {
	d, _ := os.UserConfigDir()
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
	configure(sh)
	setup(sh)
	cmd.Main(sh)
	log.Debug("end")
}

func configure(sh *ucsh.UCSh) {
	sh.Check()
	log.Debug("configure")
	cfgerr := false
	for _, fn := range cfgfiles {
		if fn == "" || fn == "ucsh.cfg" {
			continue
		}
		fh, err := os.Open(fn)
		if err != nil {
			if os.IsNotExist(err) {
				log.Debug(err)
			} else {
				cfgerr = true
				log.Error(err)
			}
		} else {
			if err := sh.Config.Load(fh.Name(), fh); err != nil {
				cfgerr = true
				log.Error(err)
			}
		}
		fh.Close()
	}
	if cfgerr {
		sh.Fail("config error")
	}
}

func setup(sh *ucsh.UCSh) {
	sh.Check()
	log.Debug("setup")
	cur, err := osuser.Current()
	if err != nil {
		log.Error(err)
		sh.Fail(err)
	}
	if err := sh.User.Load(cur); err != nil {
		log.Error(err)
		sh.Fail(err)
	}
}
