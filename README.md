vprint: classic and multicore debugging via delux print statements
======

example use (see vprint_test.go):

~~~
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
~~~

output:
~~~
vprint_test.go:13 2019-08-26 01:50:23.228 -0400 EDT hello world; a=1, b=hi
~~~

In short, vprint provides file locations, time stamps, and a
fmt.Printf() like ability to print any number of arguments
precededed by a format string.

Copyright(C) 2019 Jason E. Aten, Ph.D. All rights reserved.
License: MIT
