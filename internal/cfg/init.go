// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg

import (
	"github.com/jrmsdev/ucsh/internal/log"
)

var cinit = false
var cinitPanics = true

func Init() error {
	log.Debug("init")
	if cinit {
		if cinitPanics {
			log.Panic("config init already done")
		} else {
			log.Debug("config init alredy done... abort!")
			return nil
		}
	}
	cinit = true
	if err := load(); err != nil {
		return err
	}
	debug()
	return nil
}
