package main

import "fmt"

func main() {
	arr := []int{5, 3, 4, 6, 2}
	straightInsertionSort(arr)
	fmt.Println(arr)
}

// 直接插入排序 时间复杂度n的平方
func straightInsertionSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
