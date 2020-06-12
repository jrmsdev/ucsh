// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ucsh

import (
	"testing"
)

func TestNew(t *testing.T) {
	sh := New()
	if sh.String() != "<UCSh>" {
		t.Errorf("sh string expect: <UCSh> - got: %s", sh.String())
	}
}

func TestFail(t *testing.T) {
	sh := New()
	if sh.err != nil {
		t.Fatal("sh.err should be nil")
	}
	sh.Failf("testing.%s", "error")
	if sh.err == nil {
		t.Fatal("sh.err should not be nil")
	}
	if sh.err.Error() != "testing.error" {
		t.Errorf("sh.err message expect: testing.error - got: %s", sh.err.Error())
	}
}

func TestCheck(t *testing.T) {
	sh := New()
	if sh.err != nil {
		t.Fatal("sh.err should be nil")
	}
	defer func() {
		r := recover()
		if r != nil {
			t.Errorf("sh check should not panic: %s", r)
		}
	}()
	sh.Check()
}

func TestCheckError(t *testing.T) {
	sh := New()
	if sh.err != nil {
		t.Fatal("sh.err should be nil")
	}
	sh.Fail("testing.error")
	defer func() {
		r := recover()
		if r == nil {
			t.Error("sh check should panic")
		}
	}()
	sh.Check()
}
