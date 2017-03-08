// This file contains bindings for functions exported from Nim.

package main

// #cgo CFLAGS: -I. -Inimcache -I/shared/installed/Nim/lib
// #cgo LDFLAGS: -L. -Lnimcache libnimalgs.a -ldl
// #include <nimalgs.h>
import "C"
import "unsafe"

func initNimRuntime() {
	C.NimMain()
}

// simplest

func nim_simplest(n int) int {
	return int(C.simplest(C.int(n)))
}

// fib

func nim_rec_fib(n int) int {
	return int(C.rec_fib(C.int(n)))
}

func nim_tail_rec_fib(n int) int {
	return int(C.tail_rec_fib(C.int(n)))
}

func nim_iter_fib(n int) int {
	return int(C.iter_fib(C.int(n)))
}

func nim_mem_fib(n int) int {
	return int(C.mem_fib(C.int(n)))
}

func nim_tramp_tail_rec_fib(n int) int {
	return int(C.tramp_tail_rec_fib(C.int(n)))
}

// gcd

func nim_iter_gcd(a, b int) int {
	return int(C.iter_gcd(C.int(a), C.int(b)))
}

func nim_iter_sub_gcd(a, b int) int {
	return int(C.iter_sub_gcd(C.int(a), C.int(b)))
}

func nim_rec_gcd(a, b int) int {
	return int(C.rec_gcd(C.int(a), C.int(b)))
}

func nim_mem_gcd(a, b int) int {
	return int(C.mem_gcd(C.int(a), C.int(b)))
}

// sort

func nim_std_sort(a []int) {
	b := make([]C.int, len(a))
	for i, v := range a {
		b[i] = C.int(v)
	}
	C.std_sort((*C.int)(unsafe.Pointer(&b[0])), C.NI(len(a)))
	for i, v := range b {
		a[i] = int(v)
	}
}
