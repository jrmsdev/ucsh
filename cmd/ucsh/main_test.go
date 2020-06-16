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

	"github.com/jrmsdev/ucsh/internal/ucsh"
)

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
	userLoad(sh)
}

func TestUserLoad(t *testing.T) {
	tuser := &osuser.User{
		Uid:      "1000",
		Gid:      "1000",
		Username: "ucsht",
	}
	sh := ucsh.New()
	prevUser := osUser
	osUser = tuser
	prevErr := osUserErr
	osUserErr = nil
	defer func() {
		osUser = prevUser
		osUserErr = prevErr
	}()
	defer func() {
		r := recover()
		if r != nil {
			t.Errorf("user load should not fail: %s", r)
		}
	}()
	userLoad(sh)
}

func TestUserLoadError(t *testing.T) {
	sh := ucsh.New()
	prevUser := osUser
	osUser = &osuser.User{}
	prevErr := osUserErr
	osUserErr = nil
	defer func() {
		osUser = prevUser
		osUserErr = prevErr
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Error("user load should fail")
		}
	}()
	userLoad(sh)
}

func TestUserCfgError(t *testing.T) {
	tuser := &osuser.User{
		Uid:      "1000",
		Gid:      "1000",
		Username: "ucsht",
	}
	sh := ucsh.New()
	prevUser := osUser
	osUser = tuser
	prevErr := osUserErr
	osUserErr = nil
	prevUserCfgErr := userCfgErr
	userCfgErr = errors.New("testing.cfg.error")
	defer func() {
		osUser = prevUser
		osUserErr = prevErr
		userCfgErr = prevUserCfgErr
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Error("user cfg should fail")
		}
	}()
	userConfig(sh)
}

func TestUserConfig(t *testing.T) {
	sh := ucsh.New()
	prevUserCfgDir := userCfgDir
	userCfgDir = filepath.FromSlash("./testdata")
	prevUserCfgErr := userCfgErr
	userCfgErr = nil
	defer func() {
		userCfgDir = prevUserCfgDir
		userCfgErr = prevUserCfgErr
	}()
	defer func() {
		r := recover()
		if r != nil {
			t.Fatalf("user config should not panic: %s", r)
		}
	}()
	userConfig(sh)
}

func TestUserConfigOpenError(t *testing.T) {
	tmpdir, tmperr := ioutil.TempDir("", "ucsh_test_open_error")
	if tmperr != nil {
		t.Fatal(tmperr)
	}
	sh := ucsh.New()
	prevUserCfgDir := userCfgDir
	userCfgDir = tmpdir
	prevUserCfgErr := userCfgErr
	userCfgErr = nil
	defer func() {
		userCfgDir = prevUserCfgDir
		userCfgErr = prevUserCfgErr
		os.RemoveAll(tmpdir)
	}()
	fn := filepath.Join(tmpdir, "ucsh.cfg")
	if err := ioutil.WriteFile(fn, []byte("{}"), 0200); err != nil {
		t.Fatal(err)
	}
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("user config should panic with a permission denied error")
		}
	}()
	userConfig(sh)
}

func TestUserConfigLoadError(t *testing.T) {
	sh := ucsh.New()
	prevUserCfgDir := userCfgDir
	userCfgDir = filepath.FromSlash("./testdata/read_error")
	prevUserCfgErr := userCfgErr
	userCfgErr = nil
	defer func() {
		userCfgDir = prevUserCfgDir
		userCfgErr = prevUserCfgErr
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("user config should panic with a json parsing error")
		}
	}()
	userConfig(sh)
}
