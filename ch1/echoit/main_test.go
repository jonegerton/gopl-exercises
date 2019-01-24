package main

import (
	"testing"
)

func BenchmarkByJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ByJoin()
	}
}
func BenchmarkByLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ByLoop()
	}
}
