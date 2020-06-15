// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"testing"
)

func TestNewContainer(t *testing.T) {
	c := newContainer()
	if c.Engine != "schroot" {
		t.Errorf("c.Engine expect: schroot - got: %s", c.Engine)
	}
}
