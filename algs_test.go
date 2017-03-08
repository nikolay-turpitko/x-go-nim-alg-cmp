package main

import (
	"math/rand"
	"sort"
	"testing"
)

// High order functions used to make runSubBench more generic.
func numChk(exp int) func(b *testing.B, r interface{}) {
	return func(b *testing.B, r interface{}) {
		if exp != r.(int) {
			b.Fatalf("Expected: %d, got: %d", exp, r)
		}
	}
}

// High order functions used to make runSubBench more generic.
func wrapNumFunc(f func(int) int) func(interface{}) interface{} {
	return func(a interface{}) interface{} {
		return f(a.(int))
	}
}

// This benchmark finds cost of Nim invocation.
func BenchmarkSimplest(b *testing.B) {
	exp := 42
	for i := uint(0); i < 6; i++ {
		n := 1 << i
		runSubBench(b, "Nim", n, nil, wrapNumFunc(nim_simplest), numChk(exp))
		runSubBench(b, "Go", n, nil, wrapNumFunc(go_simplest), numChk(exp))
	}
}

func BenchmarkFib(b *testing.B) {
	exp := []int{1, 1, 3, 21, 987, 2178309}
	for i := 0; i < len(exp); i++ {
		n := 1 << uint(i)

		runSubBench(b, "GoRec", n, nil, wrapNumFunc(go_rec_fib), numChk(exp[i]))
		runSubBench(b, "GoTailRec", n, nil, wrapNumFunc(go_tail_rec_fib), numChk(exp[i]))
		runSubBench(b, "GoIter", n, nil, wrapNumFunc(go_iter_fib), numChk(exp[i]))
		runSubBench(b, "GoMem", n, nil, wrapNumFunc(go_mem_fib), numChk(exp[i]))
		runSubBench(b, "GoTramp", n, nil, wrapNumFunc(go_tramp_tail_rec_fib), numChk(exp[i]))

		runSubBench(b, "NimRec", n, nil, wrapNumFunc(nim_rec_fib), numChk(exp[i]))
		runSubBench(b, "NimTailRec", n, nil, wrapNumFunc(nim_tail_rec_fib), numChk(exp[i]))
		runSubBench(b, "NimIter", n, nil, wrapNumFunc(nim_iter_fib), numChk(exp[i]))
		runSubBench(b, "NimMem", n, nil, wrapNumFunc(nim_mem_fib), numChk(exp[i]))
		//		runSubBench(b, "NimTramp", n, nil, wrapNumFunc(nim_tramp_tail_rec_fib), numChk(exp[i]))
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

	wrapGCD := func(f func(a int, b int) int) func(interface{}) interface{} {
		return func(n interface{}) interface{} {
			i := n.(int)
			return f(t[i].a, t[i].b)
		}
	}

	for n := range t {
		runSubBench(b, "GoIter", n, nil, wrapGCD(go_iter_gcd), numChk(t[n].exp))
		runSubBench(b, "GoIterSub", n, nil, wrapGCD(go_iter_sub_gcd), numChk(t[n].exp))
		runSubBench(b, "GoRec", n, nil, wrapGCD(go_rec_gcd), numChk(t[n].exp))
		runSubBench(b, "GoMem", n, nil, wrapGCD(go_mem_gcd), numChk(t[n].exp))

		runSubBench(b, "NimIter", n, nil, wrapGCD(nim_iter_gcd), numChk(t[n].exp))
		runSubBench(b, "NimIterSub", n, nil, wrapGCD(nim_iter_sub_gcd), numChk(t[n].exp))
		runSubBench(b, "NimRec", n, nil, wrapGCD(nim_rec_gcd), numChk(t[n].exp))
		runSubBench(b, "NimMem", n, nil, wrapGCD(nim_mem_gcd), numChk(t[n].exp))
	}
}

// High order functions used to make runSubBench more generic.
func sortChk(b *testing.B, r interface{}) {
	b.StopTimer()
	defer b.StartTimer()
	if !sort.IntsAreSorted(r.([]int)) {
		b.Fatalf("Expected sorted array, but got unsorted")
	}
}

// High order functions used to make runSubBench more generic.
func wrapSortFunc(f func([]int)) func(interface{}) interface{} {
	return func(a interface{}) interface{} {
		f(a.([]int))
		return a
	}
}

func BenchmarkSort(b *testing.B) {
	const sz = 20
	const maxsz = 1 << sz
	a := make([]int, maxsz)
	for i := 0; i < maxsz; i++ {
		a[i] = rand.Int()
	}
	for i := uint(5); i < sz; i++ {
		n := 1 << i
		prepareSort := func(b *testing.B, n interface{}) interface{} {
			b.StopTimer()
			aa := make([]int, n.(int))
			copy(aa, a)
			b.StartTimer()
			return aa
		}
		runSubBench(b, "GoStd", n, prepareSort, wrapSortFunc(go_std_sort), sortChk)
		runSubBench(b, "GoNaivePar", n, prepareSort, wrapSortFunc(go_naive_par_sort), sortChk)
		runSubBench(b, "GoHeap", n, prepareSort, wrapSortFunc(go_heap_sort), sortChk)

		// Was not able to make it work.
		//runSubBench(b, "NimStd", n, prepareSort, wrapSortFunc(nim_std_sort), sortChk)
	}
}
