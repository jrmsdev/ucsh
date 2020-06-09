// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg_test

import (
	"testing"

	//~ "github.com/jrmsdev/ucsh/internal/cfg"

	"github.com/jrmsdev/ucsh/internal/_ucsht"
)

func TestInit(t *testing.T) {
	ucsht.Setup(t, "init")
}
