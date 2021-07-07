package main

import "fmt"

func main() {
	fmt.Println(fibonacci(5))
}

func fibonacci(n int) int {
	if n < 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci1(n, a1, a2 int) int {
	if n == 0 {
		return a1
	}
	return fibonacci1(n-1, a2, a1+a2)
}
