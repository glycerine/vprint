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
	vv("hello world")
}
~~~

output:
~~~
vprint_test.go:11 2019-08-26 01:42:55.695 -0400 EDT hello world
~~~

In short, vprint provides file locations, time stamps, and a
fmt.Printf() like ability to print any number of arguments
precededed by a format string.

Copyright(C) 2019 Jason E. Aten, Ph.D. All rights reserved.
License: MIT
