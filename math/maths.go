package main

// Check Even or Odd
func isEven(n int) bool {
	if (n & 1) == 0 {
		return true
	} else {
		return false
	}
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
