// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cmd

import (
	"os"
	"path/filepath"

	"github.com/jrmsdev/ucsh/internal/cmd/flags"
	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/ucsh"
)

var cfgdirs = []string{
	filepath.FromSlash("/etc/ucsh"),
	filepath.FromSlash("/usr/local/etc/ucsh"),
}

func Configure(sh *ucsh.UCSh) {
	sh.Check()
	log.Debug("configure")
	for _, dn := range cfgdirs {
		fn := filepath.Join(dn, flags.ConfigFile)
		fh, err := os.Open(fn)
		if err != nil {
			if os.IsNotExist(err) {
				log.Debug(err)
			} else {
				log.Error(err)
				fh.Close()
				sh.Fail(err)
			}
		} else {
			if err := sh.Config.Load(fh.Name(), fh); err != nil {
				fh.Close()
				sh.Fail(err)
			}
		}
		fh.Close()
	}
}
