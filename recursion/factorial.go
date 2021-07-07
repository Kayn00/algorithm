package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	fmt.Println(rescuvie(5))
	//rescuvieTail(5, 1)
	fmt.Println(time.Now().UnixNano() - start)
}

func rescuvie(n int) int {
	if n == 1 {
		return 1
	}
	return n * rescuvie(n-1)
}

func rescuvieTail(n int, a int) int {
	if n == 1 {
		return a
	}
	return rescuvieTail(n-1, n*a)
}
