package main

import "fmt"

//顺序查找
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(linearSearch1(arr, 7))
	fmt.Println(linearSearch2(arr, 7))
}

func linearSearch1(arr []int, key int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			return i
		}
	}
	return 0
}

// 去掉越界判断优化
func linearSearch2(arr []int, key int) int {
	arr = append(arr, key)
	for i := 0; ; i++ {
		if arr[i] == key {
			return i
		}
	}
}
