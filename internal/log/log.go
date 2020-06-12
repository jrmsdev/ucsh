// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	gf "fmt"
	"os"
	"runtime"
	"strconv"
)

var out = os.Stderr
var debug = false

func init() {
	env := os.Getenv("UCSH_DEBUG")
	if env != "" {
		dbg, err := strconv.ParseBool(env)
		if dbg && err == nil {
			debug = true
		}
	}
}

func Print(args ...interface{}) {
	gf.Fprintln(out, "ucsh:", gf.Sprint(args...))
}

func Printf(fmt string, args ...interface{}) {
	gf.Fprintln(out, "ucsh:", gf.Sprintf(fmt, args...))
}

func p(tag string, args ...interface{}) {
	gf.Fprintln(out, "ucsh:", tag, gf.Sprint(args...))
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
	if debug {
		p(tag("D"), args...)
	}
}

func Debugf(fmt string, args ...interface{}) {
	if debug {
		p(tag("D"), gf.Sprintf(fmt, args...))
	}
}

func Error(args ...interface{}) {
	p(tag("E"), args...)
}

func Errorf(fmt string, args ...interface{}) {
	p(tag("E"), gf.Sprintf(fmt, args...))
}
