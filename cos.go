package mylib

import "fmt"

//Factorial ...
func Factorial(n int) float64 {
	var result float64 = 1
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			result *= float64(i)
		}

	}
	return result
}

//Power ...
func Power(x float64, y int) float64 {

	var result float64 = 1
	for i := 1; i <= y; i++ {
		result *= x
	}
	return result
}

//Cos ...
func Cos(x float64) float64 {

	var result float64 = 1
	for n := 1; n <= 100; n++ {
		if (n % 2) == 0 {
			result += 1 * (Power(x, 2*n)) * (1 / Factorial(2*n))
		} else {
			result += -1 * (Power(x, 2*n)) * (1 / Factorial(2*n))
		}
	}
	return result
}
