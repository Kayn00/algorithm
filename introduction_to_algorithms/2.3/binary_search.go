package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(arr, 0, len(arr), 10))
}

func binarySearch(arr []int, begin, end, key int) int {
	if end > begin {
		mid := (begin + end) / 2
		if key == arr[mid] {
			return mid
		} else if key > arr[mid] {
			return binarySearch(arr, mid+1, end, key)
		} else {
			return binarySearch(arr, begin, mid-1, key)
		}
	}
	return -1
}
