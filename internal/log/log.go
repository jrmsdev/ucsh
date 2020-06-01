// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	gf "fmt"
	"os"
	"runtime"
)

func Print(args ...interface{}) {
	_, err := gf.Fprintln(os.Stderr, "ucsh:", gf.Sprint(args...))
	if err != nil {
		panic(err)
	}
}

func Printf(fmt string, args ...interface{}) {
	Print(gf.Sprintf(fmt, args...))
}

func p(tag string, args ...interface{}) {
	_, err := gf.Fprintln(os.Stderr, "ucsh:", tag, gf.Sprint(args...))
	if err != nil {
		panic(err)
	}
}

func dbg() string {
	tag := "[D]"
	_, fn, ln, ok := runtime.Caller(2)
	if ok {
		return gf.Sprintf("%s %s:%d", tag, fn, ln)
	}
	return tag
}

func Debug(args ...interface{}) {
	p(dbg(), args...)
}

func Debugf(fmt string, args ...interface{}) {
	p(dbg(), gf.Sprintf(fmt, args...))
}
