package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Check Even or Odd
func isEven(n int) bool {
	return (n & 1) == 0
}

// sum of n natural numbers
func findSum(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func findSum2(n int) int {
	return n * (n + 1) / 2
}

// Closest to n and Divisible by m
func closestNumber(n, m int) int {
	if m == 0 {
		return n
	}

	closest := 0
	minDifference := int(^uint(0) >> 1)

	for i := n - abs(m); i <= n+abs(m); i++ {
		if i%m == 0 {
			difference := abs(n - i)

			if difference < minDifference || (difference == minDifference && abs(i) > abs(closest)) {
				closest = i
				minDifference = difference
			}
		}
	}
	return closest
}
