package main

import (
	"fmt"
	"strings"
)

// This is where i put my logging functions. yes, its unclean, but i'll have nothing to do in this file if i put it in another file.

// prints strings or an array of it. uses `fmt.Print()` under the hood
func Print(s ...string) {
	fmt.Print(strings.Join(s, ""))
}

// same thing as Print, but with a newline at the end.
func PrintLn(s ...string) {
	fmt.Print(strings.Join(s, ""), "\n")
}
