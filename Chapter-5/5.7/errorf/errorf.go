package main

import (
	"fmt"
	"os"
)

func main() {
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
