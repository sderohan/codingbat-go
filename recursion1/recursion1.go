package recursion1

/*
Given n of 1 or more, return the factorial of n, which is n * (n-1) * (n-2) ... 1.
Compute the result recursively (without loops).

factorial(1) → 1
factorial(2) → 2
factorial(3) → 6
*/

func factorial(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return n * factorial(n-1)
}
