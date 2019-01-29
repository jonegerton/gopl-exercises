// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 12.

//!+

// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if counts[line] == nil {
				counts[line] = make(map[string]int)
			}
			counts[line][filename]++
		}
	}
	fmt.Print(counts)
	for line, files := range counts {
		for filename, n := range files {
			if len(files) > 1 || n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, filename)
			}
		}
	}
}

//!-
