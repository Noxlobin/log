package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/noxlobin/log/color"
)

// This is where i put my logging functions. yes, its unclean, but i'll have nothing to do in this file if i put it in another file.

// General Purpose Functions

// prints strings or an array of it. uses `fmt.Print()` under the hood
func Print(s ...string) {
	fmt.Print(strings.Join(s, ""))
}

// same thing as Print, but with a newline at the end.
func PrintLn(s ...string) {
	fmt.Print(strings.Join(s, ""), "\n")
}

// Error Functions

// colored error prompt.
func Error(e ...string) {
	e[0] = strings.ToLower(e[0])

	fmt.Print(color.WrapRed("error: " + strings.Join(e, "")))
	os.Exit(1)
}

// same thing as error, but with a newline.
func ErrorLn(e ...string) {
	e[0] = strings.ToLower(e[0])

	fmt.Print(color.WrapRed("error: " + strings.Join(e, "") + "\n"))
	os.Exit(1)
}

func main() {
	ErrorLn("Missing separator")
}
