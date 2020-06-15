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

var cfgfiles = []string{
	filepath.FromSlash("/etc/ucsh.cfg"),
	filepath.FromSlash("/usr/local/etc/ucsh.cfg"),
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

var osUser *osuser.User
var osUserErr error
var userCfgDir string
var userCfgErr error

func init() {
	osUser, osUserErr = osuser.Current()
	userCfgDir, userCfgErr = os.UserConfigDir()
}

func userCfgFile() (string, error) {
	if userCfgErr != nil {
		return "", userCfgErr
	}
	return filepath.Join(userCfgDir, "ucsh.cfg"), nil
}

func userLoad(sh *ucsh.UCSh) {
	sh.Check()
	log.Debug("user load")
	if osUserErr != nil {
		log.Error(osUserErr)
		sh.Fail(osUserErr)
	}
	//~ sh.User.Load(sh.Config.User, osUser)
	sh.User.Load(osUser)
}

func userConfig(sh *ucsh.UCSh) {
	sh.Check()
	log.Debug("user config")
	fn, err := userCfgFile()
	if err != nil {
		log.Error(err)
		sh.Fail(err)
	}
	cfgerr := false
	fh, err := os.Open(fn)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debug(err)
		} else {
			cfgerr = true
			log.Error(err)
		}
	} else {
		if err := sh.Config.User.Load(fn, fh); err != nil {
			cfgerr = true
			log.Error(err)
		}
	}
	if cfgerr {
		sh.Fail("user config error")
	}
}

// main

func main() {
	log.Debug("start")
	sh := ucsh.New()
	log.Debug(sh)
	configure(sh)
	userLoad(sh)
	userConfig(sh)
	cmd.Main(sh)
	log.Debug("end")
}
