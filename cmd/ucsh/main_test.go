// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"path/filepath"
	"testing"
	//~ "github.com/jrmsdev/ucsh/internal/_ucsht"
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
