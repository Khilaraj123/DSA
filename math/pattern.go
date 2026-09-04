package main

// Rectangle Pattern
func rect(rows, cols int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			print("* ")
		}
		println()
	}
}

// Program to Print Floyd's Triangle
func triangle(n int) {
	i, j, val := 1, 1, 1
	for i = 1; i <= n; i++ {
		for j = 1; j <= i; j++ {
			print(val, " ")
			val++
		}
		println()
	}
}

// Print Hollow Rectangle Pattern
func hollowRect(rows, cols int) {
	i, j := 0, 0
	for i = 1; i <= rows; i++ {
		for j = 1; j <= cols; j++ {
			if i == 1 || i == rows || j == 1 || j == cols {
				print("*")
			} else {
				print(" ")
			}
		}
		println()
	}
}
