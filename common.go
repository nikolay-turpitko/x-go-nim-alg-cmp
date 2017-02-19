package main

// Function memoize returns wrapped function which caches results on first
// execution and returns cached result afterwards.
// Note: we do not interested in general implementation here.
func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(i int) int {
		if v, ok := cache[i]; ok {
			return v
		}
		v := f(i)
		cache[i] = v
		return v
	}
}

// Function memoize2 used for same purpose as memoize, but with functions of
// 2 arguments.
func memoize2(f func(int, int) int) func(int, int) int {
	cache := make(map[int]int)
	return func(a, b int) int {
		k := a<<32 | b
		if v, ok := cache[k]; ok {
			return v
		}
		v := f(a, b)
		cache[k] = v
		return v
	}
}

// Type tramp defines a function which returns result or next trampoline.
type tramp func() (int, tramp)

// Function trampExec executes trampoline loop.
func trampoline(t tramp) int {
	res, t := t()
	for t != nil {
		res, t = t()
	}
	return res
}
