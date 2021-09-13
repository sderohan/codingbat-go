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

/*
We have bunnies standing in a line, numbered 1, 2, ... The odd bunnies (1, 3, ..) have the normal 2 ears.
The even bunnies (2, 4, ..) we'll say have 3 ears, because they each have a raised foot.
Recursively return the number of "ears" in the bunny line 1, 2, ... n (without loops or multiplication).

bunnyEars2(0) → 0
bunnyEars2(1) → 2
bunnyEars2(2) → 5
*/
func bunnyEars2(n int) int {
	if n == 0 {
		return n
	}
	if n%2 == 0 {
		return 3 + bunnyEars2(n-1)
	}
	return 2 + bunnyEars2(n-1)
}
