// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jrmsdev/ucsh/internal/log"
)

type userCfg struct {
	Shell string `json:"shell,omitempty"`
}

var User = &userCfg{
	Shell: "/bin/sh",
}

type config struct {
	User *userCfg `json:"user,omitempty"`
}

var Config = &config{
	User: User,
}

type ucsh struct {
	D *config `json:"ucsh,omitempty"`
}

var c = &ucsh{D: Config}

var cinit = false

var cfgfiles = []string{
	"/etc/ucsh.cfg",
	"/usr/local/etc/ucsh.cfg",
	userCfgFile(),
}

func userCfgFile() string {
	d, err := os.UserConfigDir()
	if err != nil {
		log.Panic(err)
	}
	return filepath.Join(d, "ucsh.cfg")
}

func load() error {
	for _, fn := range cfgfiles {
		blob, err := ioutil.ReadFile(fn)
		if err != nil {
			if os.IsNotExist(err) {
				log.Debug(err)
			} else {
				log.Error(err)
			}
		} else {
			err := json.Unmarshal(blob, c)
			if err != nil {
				log.Error(err)
				return err
			}
			log.Debugf("loaded %s", fn)
		}
	}
	return nil
}

func Init() error {
	log.Debug("init")
	if cinit {
		log.Panic("config init already done")
	}
	cinit = true
	if err := load(); err != nil {
		return err
	}
	return nil
}

func Debug() {
	log.Debugf("user.shell: %s", Config.User.Shell)
}
