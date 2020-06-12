// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var buf *bytes.Buffer

func init() {
	debug = false
	buf = new(bytes.Buffer)
	out = buf
}

func check(t *testing.T, msg, expect string) {
	t.Helper()
	got := strings.TrimSpace(buf.String())
	buf.Reset()
	x := "ucsh: " + expect
	if got != x {
		t.Errorf("%s expect: '%s' - got: '%s'", msg, x, got)
	}
}

func checkTag(t *testing.T, msg, tag, expect string) {
	t.Helper()
	got := strings.TrimSpace(buf.String())
	buf.Reset()
	prefix := fmt.Sprintf("ucsh: [%s]", tag)
	if !strings.HasPrefix(got, prefix) {
		t.Fatalf("%s prefix expect: '%s' - got: '%s'", msg, prefix, got)
	}
	if !strings.HasSuffix(got, expect) {
		t.Errorf("%s suffix expect: '%s' - got: '%s'", msg, expect, got)
	}
}

func TestPrint(t *testing.T) {
	buf.Reset()
	Print("testing")
	check(t, "Print", "testing")
	Printf("test%s", "ing")
	check(t, "Printf", "testing")
}

func TestError(t *testing.T) {
	buf.Reset()
	Error("testing")
	checkTag(t, "Error", "E", "testing")
	Errorf("test%s", "ing")
	checkTag(t, "Errorf", "E", "testing")
}

func TestDebug(t *testing.T) {
	buf.Reset()
	Debug("testing")
	if buf.String() != "" {
		t.Fatal("debug should not be enabled")
	}
	buf.Reset()
	debug = true
	defer func() {
		debug = false
	}()
	Debug("testing")
	checkTag(t, "Debug", "D", "testing")
	Debugf("test%s", "ing")
	checkTag(t, "Debugf", "D", "testing")
}

func TestPanic(t *testing.T) {
	buf.Reset()
	defer func() {
		r := recover()
		if r == nil {
			t.Error("log did not panic")
		}
		checkTag(t, "Panic", "PANIC", "testing")
	}()
	Panic("testing")
}

func TestPanicf(t *testing.T) {
	buf.Reset()
	defer func() {
		r := recover()
		if r == nil {
			t.Error("log did not panic")
		}
		checkTag(t, "Panicf", "PANIC", "testing")
	}()
	Panicf("test%s", "ing")
}
