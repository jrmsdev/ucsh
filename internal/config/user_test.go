// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewUser(t *testing.T) {
	u := newUser()
	if u.Shell != "" {
		t.Errorf("u.Shell expected: '' - got: '%s'", u.Shell)
	}
}

func TestUserLoad(t *testing.T) {
	u := newUser()
	fh, err := os.Open(filepath.FromSlash("testdata/user_load.json"))
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()
	if err := u.Load(fh.Name(), fh); err != nil {
		t.Fatal(err)
	}
	if u.Shell != filepath.FromSlash("/bin/bash") {
		t.Errorf("u.Shell expected: /bin/bash - got: %s", u.Shell)
	}
}

func TestUserLoadReadError(t *testing.T) {
	u := newUser()
	fh, err := os.Open("testdata")
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()
	err = u.Load(fh.Name(), fh)
	if err == nil {
		t.Error("user load should return a read error")
	}
}

func TestUserLoadJsonError(t *testing.T) {
	u := newUser()
	fh, err := os.Open(filepath.FromSlash("testdata/user_loaderr.json"))
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()
	err = u.Load(fh.Name(), fh)
	if err == nil {
		t.Error("user load should return a json parsing error")
	}
}
