// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"path/filepath"

	"github.com/jrmsdev/ucsh/internal/log"
)

type User struct {
	Shell string `json:"shell,omitempty"`
}

func newUser() *User {
	log.Debug("new")
	return &User{
		Shell: filepath.FromSlash("/bin/sh"),
	}
}
