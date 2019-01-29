package main

import (
	"testing"
)

var args = []string{"exe", "1", "2", "3"}

func BenchmarkEchoByJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ByJoin(args)
	}
}
func BenchmarkEchoByLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ByLoop(args)
	}
}
