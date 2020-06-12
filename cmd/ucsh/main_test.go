// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
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
