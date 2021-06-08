package main

import "fmt"

func main() {
	arr := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	ShellSort(arr)
	fmt.Println(arr)
}

// 希尔排序 其时间复杂度为n的3/2
func ShellSort(a []int) {
	n := len(a)
	h := 1
	for h < n/3 { //寻找合适的间隔h
		h = 3*h + 1
	}
	for h >= 1 {
		//将数组变为间隔h个元素有序
		for i := h; i < n; i++ {
			//间隔h插入排序
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				swap(a, j, j-h)
			}
		}
		h /= 3
		fmt.Println("h is : ", h)
	}
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
