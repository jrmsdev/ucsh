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
	return &User{}
}

func (u *User) setDefaults() {
	log.Debug("set defaults")
	u.Shell = filepath.FromSlash("/bin/sh")
}

func (u *User) kmap() map[string]*string {
	return map[string]*string{
		"shell": &u.Shell,
	}
}

type userConfig struct {
	User *User `json:"user,omitempty"`
}

type ucshUser struct {
	D *userConfig `json:"ucsh,omitempty"`
}

func (u *User) Load(fn string, fh io.Reader) error {
	log.Debugf("load %s", fn)
	blob, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Error(err)
		return err
	}
	obj := &ucshUser{D: &userConfig{User: u}}
	if err := json.Unmarshal(blob, obj); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
