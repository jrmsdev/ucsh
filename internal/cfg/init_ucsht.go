// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// +build ucsht

package cfg

import (
	"testing"

	//~ "github.com/jrmsdev/ucsh/internal/log"
)

func InitTest(t *testing.T) {
	t.Helper()
	cinitPanics = false
	err := Init()
	if err != nil {
		t.Fatal(err)
	}
}

func InitTestCleanup(t *testing.T) {
	t.Helper()
	cinitPanics = true
}
