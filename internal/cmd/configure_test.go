// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/jrmsdev/ucsh/internal/ucsh"
)

var cfgfilesOrig = cfgfiles

func TestConfigure(t *testing.T) {
	sh := ucsh.New()
	cfgfiles = []string{filepath.FromSlash("testdata/config.json")}
	defer func() {
		cfgfiles = cfgfilesOrig
	}()
	defer func() {
		r := recover()
		if r != nil {
			t.Fatalf("configure should not fail: %s", r)
		}
	}()
	Configure(sh)
}

func TestCfgNotFound(t *testing.T) {
	sh := ucsh.New()
	cfgfiles = []string{filepath.FromSlash("testdata/nofile.cfg")}
	defer func() {
		cfgfiles = cfgfilesOrig
	}()
	defer func() {
		r := recover()
		if r != nil {
			t.Fatalf("configure should not fail if file does not exists: %s", r)
		}
	}()
	Configure(sh)
}

func TestCfgOpenError(t *testing.T) {
	sh := ucsh.New()
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
		cfgfiles = cfgfilesOrig
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Error("configure should fail with a permission denied error")
		}
	}()
	Configure(sh)
}

func TestCfgError(t *testing.T) {
	sh := ucsh.New()
	cfgfiles = []string{filepath.FromSlash("./testdata")}
	defer func() {
		cfgfiles = cfgfilesOrig
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Error("configure should fail trying to read a directory")
		}
	}()
	Configure(sh)
}
