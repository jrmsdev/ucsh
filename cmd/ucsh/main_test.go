// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"errors"
	"io/ioutil"
	"os"
	osuser "os/user"
	"path/filepath"
	"testing"

	"github.com/jrmsdev/ucsh"
)

var tfiles = []string{
	filepath.Join("", "ucsh.cfg"),
	filepath.FromSlash("testdata/ucsh.cfg"),
}

func TestMain(t *testing.T) {
	cfgfiles = tfiles
	main()
}

func TestCfgNotFound(t *testing.T) {
	cfgfiles = []string{filepath.FromSlash("testdata/nofile.cfg")}
	defer func() {
		r := recover()
		if r != nil {
			t.Errorf("config not found failed: %s", r)
		}
	}()
	main()
}

func TestCfgError(t *testing.T) {
	cfgfiles = []string{"testdata"}
	defer func() {
		r := recover()
		if r == nil {
			t.Error("config error did not fail")
		}
	}()
	main()
}

func TestOpenError(t *testing.T) {
	fh, err := ioutil.TempFile("", "ucsh.test.read.err.")
	if err != nil {
		t.Fatal(err)
	}
	tmpfn := fh.Name()
	fh.Close()
	defer os.Remove(tmpfn)
	if err := os.Chmod(tmpfn, 0200); err != nil {
		t.Fatal(err)
	}
	cfgfiles = []string{tmpfn}
	defer func() {
		r := recover()
		if r == nil {
			t.Error("open error did not fail")
		}
	}()
	main()
}

func TestOsUserError(t *testing.T) {
	sh := ucsh.New()
	prevErr := osUserErr
	osUserErr = errors.New("testing os user error")
	defer func() {
		osUserErr = prevErr
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Error("os user error did not fail")
		}
	}()
	setup(sh)
}

func TestUserLoad(t *testing.T) {
	tuser := &osuser.User{
		Uid: "1000",
		Gid: "1000",
		Username: "ucsht",
	}
	sh := ucsh.New()
	prevUser := osUser
	osUser = tuser
	prevErr := osUserErr
	osUserErr = nil
	defer func() {
		osUser = prevUser
	}()
	defer func() {
		osUserErr = prevErr
	}()
	defer func() {
		r := recover()
		if r != nil {
			t.Errorf("user load should not fail: %s", r)
		}
	}()
	setup(sh)
}

func TestUserLoadError(t *testing.T) {
	sh := ucsh.New()
	prevUser := osUser
	osUser = &osuser.User{}
	prevErr := osUserErr
	osUserErr = nil
	defer func() {
		osUser = prevUser
	}()
	defer func() {
		osUserErr = prevErr
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Error("user load should fail")
		}
	}()
	setup(sh)
}
