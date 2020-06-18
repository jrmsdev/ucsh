// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cmd

import (
	"os"
	osuser "os/user"
	"path/filepath"

	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/ucsh"
)

var osUser *osuser.User
var osUserErr error
var userCfgDir string
var userCfgErr error

func init() {
	osUser, osUserErr = osuser.Current()
	userCfgDir, userCfgErr = os.UserConfigDir()
	if userCfgDir != "" {
		userCfgDir = filepath.Join(userCfgDir, "ucsh")
	}
}

func UserLoad(sh *ucsh.UCSh) {
	sh.Check()
	log.Debug("user load")
	if osUserErr != nil {
		log.Error(osUserErr)
		sh.Fail(osUserErr)
	}
	sh.User.Load(osUser)
}

func userCfgFile() (string, error) {
	if userCfgErr != nil {
		return "", userCfgErr
	}
	return filepath.Join(userCfgDir, "config.json"), nil
}

func UserConfig(sh *ucsh.UCSh) {
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
