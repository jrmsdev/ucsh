// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ucsh

import (
	"context"
	"errors"
	"fmt"

	"github.com/jrmsdev/ucsh/internal/config"
	"github.com/jrmsdev/ucsh/internal/log"
)

var bgctx context.Context

func init() {
	bgctx = context.Background()
}

type UCSh struct {
	ctx    context.Context
	cancel context.CancelFunc
	err    error
	Config *config.Config
}

func New() *UCSh {
	log.Debug("new")
	ctx, cancel := context.WithCancel(bgctx)
	return &UCSh{
		ctx:    ctx,
		cancel: cancel,
		Config: config.New(),
	}
}

func (sh *UCSh) String() string {
	return "<UCSh>"
}

func (sh *UCSh) Fail(args ...interface{}) {
	log.Debug("fail")
	sh.err = errors.New(fmt.Sprint(args...))
	sh.cancel()
}

func (sh *UCSh) Failf(f string, args ...interface{}) {
	log.Debug("failf")
	sh.err = errors.New(fmt.Sprintf(f, args...))
	sh.cancel()
}

func (sh *UCSh) Check() {
	log.Debug("check context")
	if err := sh.ctx.Err(); err != nil {
		if sh.err != nil {
			log.Errorf("context failed: %s", sh.err)
		}
		log.Panic(err)
	}
}
