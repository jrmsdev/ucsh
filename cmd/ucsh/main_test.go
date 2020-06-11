// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"path/filepath"
	"testing"

	//~ "github.com/jrmsdev/ucsh/internal/_ucsht"
)

func init() {
	cfgfiles = []string{
		filepath.FromSlash("testdata/ucsh.cfg"),
	}
}

func TestMain(t *testing.T) {
	main()
}
