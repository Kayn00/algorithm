package main

import "fmt"

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
