// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 12.

//!+

// Dup prints the count and text of lines that
// appear more than once in the named input files.
// Includes the name of the file in the output
// To run use command similar to:
// go run main.go "C:\Git\gopl-exercises\ch1\dup\test1.txt" "C:\Git\gopl-exercises\ch1\dup\test2.txt" "C:\Git\gopl-exercises\ch1\dup\test3.txt"
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

			//Allow for windows line breaks
			line = strings.Replace(line, "\r", "", -1)

			if len(line) == 0 {
				continue
			}

			if counts[line] == nil {
				counts[line] = make(map[string]int)
			}
			counts[line][filename]++
		}
	}
	//fmt.Print(counts)
	for line, files := range counts {
		for filename, n := range files {
			if len(files) > 1 || n > 1 {
				fmt.Printf("%s\t%s\t%d\n", line, filename, n)
			}
		}
	}
}

//!-
