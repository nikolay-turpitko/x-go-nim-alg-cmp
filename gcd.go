package main

func go_iter_gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func go_iter_sub_gcd(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func go_rec_gcd(a, b int) int {
	if b != 0 {
		a = go_rec_gcd(b, a%b)
	}
	return a
}

var go_mem_gcd = memoize2(go_rec_gcd)
