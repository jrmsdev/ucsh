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

func tag(s string) string {
	t := gf.Sprintf("[%s]", s)
	_, fn, ln, ok := runtime.Caller(2)
	if ok {
		return gf.Sprintf("%s %s:%d", t, fn, ln)
	}
	return t
}

func Panic(args ...interface{}) {
	p(tag("PANIC"), args...)
	panic("ucsh")
}

func Debug(args ...interface{}) {
	p(tag("D"), args...)
}

func Debugf(fmt string, args ...interface{}) {
	p(tag("D"), gf.Sprintf(fmt, args...))
}

func Error(args ...interface{}) {
	p(tag("E"), args...)
}

func Errorf(fmt string, args ...interface{}) {
	p(tag("E"), gf.Sprintf(fmt, args...))
}
