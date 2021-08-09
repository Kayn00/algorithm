package main

import "fmt"

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	mergeSort(arr, 0, len(arr))
	fmt.Println(arr)
}

func mergeSort(arr []int, min, max int) {
	if max-min > 1 {
		mid := (min + max) / 2
		mergeSort(arr, min, mid)
		mergeSort(arr, mid, max)
		merge(arr, min, mid, max)
	}
}

func merge(arr []int, min, mid, max int) {
	leftSize := mid - min
	rightSize := max - mid
	leftIndex, rightIndex := 0, 0
	result := make([]int, 0, max-min)
	for leftIndex < leftSize && rightIndex < rightSize {
		leftValue := arr[min+leftIndex]
		rightValue := arr[mid+rightIndex]
		if leftValue < rightValue {
			result = append(result, leftValue)
			leftIndex += 1
		} else {
			result = append(result, rightValue)
			rightIndex += 1
		}
	}

	result = append(result, arr[min+leftIndex:mid]...)
	result = append(result, arr[mid+rightIndex:max]...)

	for index, _ := range result {
		arr[min+index] = result[index]
	}
}
