// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ucsht

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/jrmsdev/ucsh/internal/cfg"
)

type ucshTest struct {
	ctx *context.Context
	cancel *context.CancelFunc
	t *testing.T
	conf string
}

func newTest(cur *ucshTest, parent *testing.T) {
	parent.Log("new")
	cur.ctx = nil
	cur.cancel = nil
	ctx, cancel := context.WithCancel(context.Background())
	cur.ctx = &ctx
	cur.cancel = &cancel
	cur.t = parent
	cur.conf = ""
}

var cur = new(ucshTest)

func Setup(t *testing.T, src string) {
	newTest(cur, t)
	cur.conf = filepath.Join("testdata", src + ".json")
	cur.t.Logf("setup: %s", cur.conf)
	cfg.InitTest(cur.t)
}

func Cleanup() {
	cur.t.Logf("cleanup: %s", cur.conf)
	cfg.InitTestCleanup(cur.t)
}
