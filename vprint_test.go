package vprint_test

import (
	"github.com/glycerine/vprint"
	"testing"
)

var vv = vprint.VV

func TestHello(t *testing.T) {
	vv("hello world")
}
