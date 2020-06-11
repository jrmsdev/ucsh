// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"os"
	"path/filepath"

	"github.com/jrmsdev/ucsh/internal/log"
)

func userCfgFile() string {
	d, err := os.UserConfigDir()
	if err != nil {
		log.Panic(err)
	}
	return filepath.Join(d, "ucsh.cfg")
}

type User struct {
	Shell string `json:"shell,omitempty"`
}

func newUser() *User {
	return &User{
		Shell: filepath.FromSlash("/bin/sh"),
	}
}

func (u *User) debug() {
	log.Debugf("user.shell: %s", u.Shell)
}
