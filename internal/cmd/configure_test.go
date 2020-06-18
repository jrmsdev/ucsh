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

var cfgdirsOrig = cfgdirs

func TestConfigure(t *testing.T) {
	sh := ucsh.New()
	cfgdirs = []string{filepath.FromSlash("testdata")}
	defer func() {
		cfgdirs = cfgdirsOrig
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
	cfgdirs = []string{filepath.FromSlash("testdata/nodir")}
	defer func() {
		cfgdirs = cfgdirsOrig
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
	cfgdirs = []string{tmpfn}
	defer func() {
		cfgdirs = cfgdirsOrig
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
	cfgdirs = []string{filepath.FromSlash("./testdata/read_error")}
	defer func() {
		cfgdirs = cfgdirsOrig
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Error("configure should fail trying to read a directory")
		}
	}()
	Configure(sh)
}
