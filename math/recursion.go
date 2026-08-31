package main

import "fmt"

// Sum of Natural Numbers
func sum(n int) int {
	if n == 1 {
		return 1
	}
	return n + sum(n-1)
}

// Factorial
func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}

// demonstrate working of recursion
func printFunc(test int) int {
	if test < 1 {
		return 1
	} else {
		fmt.Print(" ", test)
		printFunc(test - 1)
		print(" ", test)
		return 0
	}
}

// Fibonacci with Recursion
func fibonacci(num int) int{
	if num == 0 {
		return 0
	} else if num == 1 || num == 2{
		return 1
	} else {
		return (fibonacci(num-1)+ fibonacci(num-2))
	}
}

func main() {
	n := 5
	fmt.Println("The sum of numbers are", sum(n))
	fmt.Println("The factorial of numbers are", fact(n))
	printFunc(n)
	println(" ")
	print("Fibonacci series of number 5 is: ")
	for i := 0; i < n; i++ {
		print(fibonacci(i), " ")
	}
}
