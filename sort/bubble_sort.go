package main

import "fmt"

// 冒泡排序 时间复杂度n的平方
func main() {
	arr := []int{7, 1, 9, 6, 5, 4, 2, 3, 8}
	bubbleSort(arr)
	fmt.Println(arr)

	//arr = []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	//bubbleSort1(arr)
	//fmt.Println(arr)

	arr = []int{7, 1, 9, 6, 5, 4, 2, 3, 8}
	bubbleSort2(arr)
	fmt.Println(arr)
}

// 正向排序，每次循环对后面没有影响
func bubbleSort(arr []int) {
	compare := 0
	swap := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			compare++
			if arr[j] < arr[i] {
				swap++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		fmt.Println(i, "compare is : ", compare)
		fmt.Println(i, "swap is : ", swap)
	}
	fmt.Println("compare is : ", compare)
	fmt.Println("swap is : ", swap)
}

// 反向排序，每次循环后面的数字变换会简单
func bubbleSort1(arr []int) {
	compare := 0
	swap := 0
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 2; j >= i; j-- {
			compare++
			if arr[j] > arr[j+1] {
				swap++
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(i, "compare is : ", compare)
		fmt.Println(i, "swap is : ", swap)
	}
	fmt.Println("compare is : ", compare)
	fmt.Println("swap is : ", swap)
}

// 添加flag标示，当次循环没有数据交换，退出后续循环
func bubbleSort2(arr []int) {
	compare := 0
	swap := 0
	flag := true
	for i := 0; i < len(arr) && flag; i++ {
		flag = false
		for j := len(arr) - 2; j >= i; j-- {
			compare++
			if arr[j] > arr[j+1] {
				swap++
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true // 如果有数据交换，则flag为true
			}
		}
		fmt.Println("compare is : ", compare)
		fmt.Println("swap is : ", swap)
	}
	fmt.Println("compare is : ", compare)
	fmt.Println("swap is : ", swap)
}
