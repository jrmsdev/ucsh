// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"testing"
)

func TestNew(t *testing.T) {
	cfg, err := New()
	t.Log(cfg, err)
}
