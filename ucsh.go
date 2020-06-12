// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ucsh

import (
	"context"
	"errors"
	"fmt"

	"github.com/jrmsdev/ucsh/internal/config"
	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/user"
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
	User   *user.User
}

func New() *UCSh {
	log.Debug("new")
	ctx, cancel := context.WithCancel(bgctx)
	return &UCSh{
		ctx:    ctx,
		cancel: cancel,
		Config: config.New(),
		User:   user.New(),
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
	sh.Fail(fmt.Sprintf(f, args...))
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
