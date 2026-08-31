# Recursion
- Sum of natural numbers
```go
func sum(n int) int {
	if n == 1 {
		return 1
	}
	return n + sum(n-1)
}
```
- Factorial
```go
func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}
```
- Recursive CallStack Behavior
```go
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
```
- Fibonacci Sequence
```go
func fibonacci(num int) int{
	if num == 0 {
		return 0
	} else if num == 1 || num == 2{
		return 1
	} else {
		return (fibonacci(num-1)+ fibonacci(num-2))
	}
}
```

# Maths
- Check Even or Odd
```go
func isEven(n int) bool {
	if (n & 1) == 0 {
		return true
	} else {
		return false
	}
}
```
- sum of n natural numbers
```go
func findSum(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}
```
- sum of n natural numbers (method 2)
```go
func findSum2(n int) int {
	return n * (n + 1) / 2
}
```