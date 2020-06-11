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
	if err := c.Load("load.json", fh); err != nil {
		t.Error(err)
	}
	if c.User.Shell != filepath.FromSlash("/bin/bash") {
		t.Errorf("User.Shell expected: /bin/bash - got: %s", c.User.Shell)
	}
}

func TestLoadError(t *testing.T) {
	c := New()
	fh, err := os.Open(filepath.FromSlash("testdata/loaderr.json"))
	if err != nil {
		t.Fatal(err)
	}
	fh.Close()
	if err := c.Load("loaderr.json", fh); err == nil {
		t.Error("load error did not fail")
	} else if err.Error() != "read testdata/loaderr.json: file already closed" {
		t.Errorf("load json invalid error: %s", err.Error())
	}
}

func TestLoadJsonError(t *testing.T) {
	c := New()
	fh, err := os.Open(filepath.FromSlash("testdata/loaderr.json"))
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()
	if err := c.Load("loaderr.json", fh); err == nil {
		t.Error("load json error did not fail")
	} else if err.Error() != "unexpected end of JSON input" {
		t.Errorf("load json invalid error: %s", err.Error())
	}
}
