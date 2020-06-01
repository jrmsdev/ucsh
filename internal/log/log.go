// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	gf "fmt"
	"io"
	"os"
	"runtime"
)

func Print(args ...interface{}) {
	_, err := gf.Fprintln(os.Stderr, args...)
	if err != nil {
		panic(err)
	}
}

func Printf(fmt string, args ...interface{}) {
	_, err := gf.Fprintln(os.Stderr, gf.Sprintf(fmt, args...))
	if err != nil {
		panic(err)
	}
}

func p(out io.Writer, tag string, fmt string, args ...interface{}) {
	_, err := gf.Fprintln(out, tag, gf.Sprintf(fmt, args...))
	if err != nil {
		panic(err)
	}
}

func Debug(fmt string, args ...interface{}) {
	tag := "D:"
	_, fn, ln, ok := runtime.Caller(1)
	if ok {
		tag = gf.Sprintf("D: %s:%d", fn, ln)
	}
	p(os.Stderr, tag, fmt, args...)
}
