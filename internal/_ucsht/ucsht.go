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

var bgctx = context.Background()

func newTest(cur *ucshTest, parent *testing.T, src string) {
	parent.Log("new")
	cur.ctx = nil
	cur.cancel = nil
	ctx, cancel := context.WithCancel(bgctx)
	cur.ctx = &ctx
	cur.cancel = &cancel
	cur.t = nil
	cur.t = parent
	cur.conf = filepath.Join("testdata", src + ".json")
}

var cur = new(ucshTest)

func cleanup(t *testing.T, conf string) {
	t.Logf("cleanup: %s", conf)
	cfg.InitTestCleanup(t, conf)
}

func Setup(t *testing.T, src string) {
	newTest(cur, t, src)
	cleanup(cur.t, cur.conf)
	cur.t.Logf("setup: %s", cur.conf)
	cfg.InitTest(cur.t, cur.conf)
}
