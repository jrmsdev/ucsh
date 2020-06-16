// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"testing"
)

func init() {
	args = []string{}
}

func TestMain(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil {
			t.Fatalf("main should not fail: %s", r)
		}
	}()
	main()
}
