package main

import "unsafe"

/*
#include <stdio.h>
#include <stdlib.h>
void hello(char *s) { printf("Hello, %s\n", s); }
*/
import "C"

// Comments before import "C" would be built as C codes
func main() {
	str := C.CString("Athens Gophers")
	C.hello(str) // Call c function from Go code
	C.free(unsafe.Pointer(str))
}
