package main

func go_rec_fib(n int) int {
	if n <= 2 {
		return 1
	}
	return go_rec_fib(n-1) + go_rec_fib(n-2)
}

func go_tail_rec_fib(n int) int {
	var f func(int, int, int) int
	f = func(term, val, prev int) int {
		if term == 0 {
			return prev
		}
		if term == 1 {
			return val
		}
		return f(term-1, prev+val, val)
	}
	return f(n, 1, 0)
}

func go_iter_fib(n int) int {
	if n <= 2 {
		return 1
	}
	prev1 := 0
	prev2 := 1
	for i := 2; i <= n; i++ {
		prev1, prev2 = prev2, prev1+prev2
	}
	return prev2
}

var go_mem_fib = memoize(go_rec_fib)

func tramp_tail_rec_fib(term, val, prev int) tramp {
	if term == 0 {
		return func() (int, tramp) {
			return prev, nil
		}
	}
	if term == 1 {
		return func() (int, tramp) {
			return val, nil
		}
	}
	return func() (int, tramp) {
		return 0, tramp_tail_rec_fib(term-1, prev+val, val)
	}
}

func go_tramp_tail_rec_fib(n int) int {
	return trampoline(tramp_tail_rec_fib(n, 1, 0))
}
