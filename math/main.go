package main

import "fmt"

func main() {
	n := 10

	//From recursion.go
	fmt.Println("The sum of numbers are", sum(n))
	fmt.Println("The factorial of numbers are", fact(n))
	printFunc(n)
	println(" ")
	print("Fibonacci series of number 5 is: ")
	for i := 0; i < n; i++ {
		print(fibonacci(i), " ")
	}

	//From maths.go
	if isEven(n) {
		println("EvenNumber")
	} else {
		println("Odd Number")
	}

	println("The sum is: ", findSum(n))
	println("The sum by another method is: ", findSum2(n))

	a, b := 10, 15
	println("The closest number to", a, "and divisible by", b, "is:", closestNumber(a, b))

	if isSumofConsecutiveNumbers(n) {
		println("true")
	} else {
		println("false")
	}
}
