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

// runSubBench performs individual benchmark.
// name - name of the test
// n - size of the test (like N in big-O notation)
// prep - func to prepare data (can be nil)
// test - test itself
// chk - func to check result (can be nil)
func runSubBench(
	b *testing.B,
	name string,
	n int,
	prep func(b *testing.B, n interface{}) interface{},
	test func(n interface{}) interface{},
	chk func(b *testing.B, r interface{})) {
	b.Run(
		fmt.Sprintf("%s#%d", name, n),
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var d interface{}
				if prep == nil {
					d = n
				} else {
					d = prep(b, n)
				}
				r := test(d)
				if chk != nil {
					chk(b, r)
				}
			}
		})
}
