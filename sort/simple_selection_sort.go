package main

import "fmt"

// 简单选择排序 时间复杂度n的平方
func main() {
	arr := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	simpleSelectionSort(arr)
	fmt.Println(arr)
}

func simpleSelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
