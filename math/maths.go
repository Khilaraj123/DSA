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

// Express a number as sum of consecutive numbers
// Given a positive integer n,
// find whether it can be represented as the sum of two or more consecutive positive integers.
func isSumofConsecutiveNumbers(n int) bool {
	if n <= 0 {
		return false
	}

	for start := 1; start < n; start++ {
		sum := 0

		for end := start; sum < n; end++ {
			sum += end

			if sum == n && end > 1 {
				return true
			}

			if sum > n {
				break
			}
		}
	}
	return false
}
