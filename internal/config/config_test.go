// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNew(t *testing.T) {
	c := New()
	if c.User.Shell != filepath.FromSlash("/bin/sh") {
		t.Errorf("User.Shell expected: /bin/sh - got: %s", c.User.Shell)
	}
}

func TestLoad(t *testing.T) {
	c := New()
	fh, err := os.Open(filepath.FromSlash("testdata/load.json"))
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()
	if err := c.Load(fh); err != nil {
		t.Error(err)
	}
}
