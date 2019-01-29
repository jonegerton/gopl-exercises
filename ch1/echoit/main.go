// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//!+
func main() {

	fmt.Println(ByJoin(os.Args))

	fmt.Println(ByLoop(os.Args))
}

func ByJoin(args []string) string {
	return strings.Join(os.Args[1:], " ")
}

func ByLoop(args []string) string {
	var builder strings.Builder
	for i, arg := range os.Args[1:] {
		builder.WriteString(strconv.Itoa(i) + ":" + arg + "\n")
	}
	return builder.String()
}

//!-
