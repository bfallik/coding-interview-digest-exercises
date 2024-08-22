package ex10

func Factorial(n int) int {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}

func Combinations(n, r int) int {
	return Factorial(n) / (Factorial(n-r) * Factorial(r))
}

func Solve(m, n int) int {
	return Combinations(m+n-2, n-1)
}
