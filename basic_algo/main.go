package main

import "fmt"

func main() {
	fmt.Printf("** Algorithms **\n")

	// Fibonacci
	fib6, _ := fibonacci(6)
	fib10, _ := fibonacci(10)
	_, err := fibonacci(25000)
	fmt.Printf("Fibo of 6 : %v\n", fib6)
	fmt.Printf("Fibo of 10 : %v\n", fib10)
	if err != nil {
		fmt.Println(err)
	}

	// Factorial
	fact6, _ := factorial(6)
	fact10, _ := factorial(10)
	_, err = factorial(25000)
	fmt.Printf("Factorial of 6 : %v\n", fact6)
	fmt.Printf("Factorial of 10 : %v\n", fact10)
	if err != nil {
		fmt.Println(err)
	}
}