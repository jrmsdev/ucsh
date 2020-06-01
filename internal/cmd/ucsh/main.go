// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ucsh

import (
	"fmt"
	"os"

	"github.com/jrmsdev/ucsh/internal/env"
	"github.com/jrmsdev/ucsh/internal/log"
	"github.com/jrmsdev/ucsh/internal/user"
)

func Main() {
	log.Debug("main start")
	e := env.New()
	u, err := user.New(e)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
	log.Debug("main end")
}
