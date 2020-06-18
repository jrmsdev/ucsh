// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"testing"
)

func TestNewContainer(t *testing.T) {
	c := newContainer()
	if c.Engine != "" {
		t.Errorf("c.Engine expect: '' - got: '%s'", c.Engine)
	}
}
