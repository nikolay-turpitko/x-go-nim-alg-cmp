package main

import (
	"testing"
)

// This benchmark finds cost of Nim invocation.
func BenchmarkSimplest(b *testing.B) {
	exp := 42
	for i := uint(0); i < 6; i++ {
		n := 1 << i
		runSubBench(b, "Nim", n, exp, nim_simplest)
		runSubBench(b, "Go", n, exp, go_simplest)
	}
}

func BenchmarkFib(b *testing.B) {
	exp := []int{1, 1, 3, 21, 987, 2178309}
	for i := 0; i < len(exp); i++ {
		n := 1 << uint(i)

		runSubBench(b, "GoRec", n, exp[i], go_rec_fib)
		runSubBench(b, "GoTailRec", n, exp[i], go_tail_rec_fib)
		runSubBench(b, "GoIter", n, exp[i], go_iter_fib)
		runSubBench(b, "GoMem", n, exp[i], go_mem_fib)
		runSubBench(b, "GoTramp", n, exp[i], go_tramp_tail_rec_fib)

		runSubBench(b, "NimRec", n, exp[i], nim_rec_fib)
		runSubBench(b, "NimTailRec", n, exp[i], nim_tail_rec_fib)
		runSubBench(b, "NimIter", n, exp[i], nim_iter_fib)
		runSubBench(b, "NimMem", n, exp[i], nim_mem_fib)
		//		runSubBench(b, "NimTramp", n, exp[i], nim_tramp_tail_rec_fib)
	}
}

func BenchmarkGCD(b *testing.B) {
	t := map[int]struct {
		a   int
		b   int
		exp int
	}{
		// Fibonacci N and N+1 are used to build series (theoretical worst case).
		1:  {1, 1, 1},
		2:  {1, 2, 1},
		4:  {3, 5, 1},
		8:  {21, 34, 1},
		16: {987, 1597, 1},
		24: {121393, 75025, 1},
		32: {2178309, 3524578, 1},
	}
	for n, _ := range t {
		runSubBench(b, "GoIter", n, t[n].exp, func(n int) int { return go_iter_gcd(t[n].a, t[n].b) })
		runSubBench(b, "GoIterSub", n, t[n].exp, func(n int) int { return go_iter_sub_gcd(t[n].a, t[n].b) })
		runSubBench(b, "GoRec", n, t[n].exp, func(n int) int { return go_rec_gcd(t[n].a, t[n].b) })
		runSubBench(b, "GoMem", n, t[n].exp, func(n int) int { return go_mem_gcd(t[n].a, t[n].b) })

		runSubBench(b, "NimIter", n, t[n].exp, func(n int) int { return nim_iter_gcd(t[n].a, t[n].b) })
		runSubBench(b, "NimIterSub", n, t[n].exp, func(n int) int { return nim_iter_sub_gcd(t[n].a, t[n].b) })
		runSubBench(b, "NimRec", n, t[n].exp, func(n int) int { return nim_rec_gcd(t[n].a, t[n].b) })
		runSubBench(b, "NimMem", n, t[n].exp, func(n int) int { return nim_mem_gcd(t[n].a, t[n].b) })
	}
}
