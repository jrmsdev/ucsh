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

var cinit = false
var cinitPanics = true

var cfgfiles = []string{
	filepath.FromSlash("/etc/ucsh.cfg"),
	filepath.FromSlash("/usr/local/etc/ucsh.cfg"),
	userCfgFile(),
}

func userCfgFile() string {
	d, err := os.UserConfigDir()
	if err != nil {
		log.Panic(err)
	}
	return filepath.Join(d, "ucsh.cfg")
}

func Init() error {
	log.Debug("init")
	if cinit {
		if cinitPanics {
			log.Panic("config init already done")
		} else {
			log.Debug("config init alredy done... abort!")
		}
	} else {
		cinit = true
		for _, fn := range cfgfiles {
			if err := load(fn); err != nil {
				return err
			}
		}
	}
	debug()
	return nil
}

func load(fn string) error {
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
	return nil
}
