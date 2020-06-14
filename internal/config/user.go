// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
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

func (u *User) Load(name string, fh io.Reader) error {
	log.Debugf("load %s", name)
	blob, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Error(err)
		return err
	}
	if err := json.Unmarshal(blob, u); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
