// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"testing"

	"github.com/jrmsdev/ucsh/internal/_ucsht"
)

func TestMain(t *testing.T) {
	ucsht.Setup(t, "main")
	main()
}
