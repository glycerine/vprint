package vprint_test

import (
	"github.com/glycerine/vprint"
	"testing"
)

var vv = vprint.VV

func TestHello(t *testing.T) {
	a := 1
	b := "hi"
	vv("hello world; a=%v, b=%v", a, b)
}
