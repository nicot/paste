package main

import (
	"testing"
)

func TestGetName(t *testing.T) {
	v := getName(0)
	if v != "a" {
		t.Error("zero should give a. Got %v", v)
	}
	v = getName(20)
	if v != "u" {
		t.Error("zero should give u. Got %v", v)
	}
	v = getName(25)
	if v != "ja" {
		t.Errorf("200 should give something like ja %v", v)
	}
	v = getName(300)
	if v != "ja" {
		t.Errorf("200 should give something like ja %v", v)
	}
}
