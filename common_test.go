package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	initNimRuntime()
	os.Exit(m.Run())
}

func runSubBench(
	b *testing.B,
	name string, // Name of test
	n int, // "Size" of test (like N in big-O notation)
	exp int, // Expected result
	f func(n int) int) {
	b.Run(
		fmt.Sprintf("%s#%d", name, n),
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r := f(n)
				if exp != r {
					b.Fatalf("Expected: %d, got: %d", exp, r)
				}
			}
		})
}
