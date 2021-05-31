package main

import "fmt"

func main() {
	arr := []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}
	//fmt.Println(binarySearch(arr, 99))
	//fmt.Println(interpolationSearch(arr, 99))

	fmt.Println(fibonacciSearch(arr, 99))
}

// 折半查找
func binarySearch(arr []int, key int) (bool, int, int) {
	if len(arr) == 0 {
		return false, 0, 0
	}

	var low, high, mid, n int
	low = 0
	high = len(arr) - 1

	for low <= high {
		n++
		mid = (low + high) / 2
		if key < arr[mid] {
			high = mid - 1
		} else if key > arr[mid] {
			low = mid + 1
		} else {
			return true, mid, n
		}
	}
	return false, 0, n
}

// 插值查找
// mid=(low+high)/2=low+(high-low)/2
// mid=low+(high-low)*(key-a[low])/(a[high]-a[low])
func interpolationSearch(arr []int, key int) (bool, int, int) {
	if len(arr) == 0 {
		return false, 0, 0
	}

	var low, high, mid, n int
	low = 0
	high = len(arr) - 1

	for low <= high {
		n++
		mid = low + (high-low)*(key-arr[low])/(arr[high]-arr[low])
		fmt.Println("mid is : ", mid)
		if key < arr[mid] {
			high = mid - 1
		} else if key > arr[mid] {
			low = mid + 1
		} else {
			return true, mid, n
		}
	}
	return false, 0, n
}

// 斐波那契查找
func fibonacciSearch(arr []int, key int) int {
	var low, high, mid, maxIndex, k int
	low = 0
	maxIndex = len(arr) - 1
	high = maxIndex
	k = 0
	//fmt.Println("arr cap is : ", cap(arr))

	// 计算n位于斐波那契数列的位置
	for maxIndex > fibonacci(k)-1 {
		k++
	}

	// 将不满的数值补全
	fmt.Println("fibonacci(k)-1 is : ", fibonacci(k)-1)
	for i := maxIndex; i < fibonacci(k)-1; i++ {
		arr = append(arr, arr[maxIndex])
	}
	fmt.Println("arr is : ", arr)

	for low <= high {
		mid = low + fibonacci(k-1) - 1
		if key < arr[mid] {
			high = mid - 1
			k = k - 1
		} else if key > arr[mid] {
			low = mid + 1
			k = k - 2
		} else {
			if mid <= maxIndex {
				return mid
			} else {
				return maxIndex
			}
		}
	}
	return len(arr)
}

func fibonacci(k int) int {
	if k < 3 {
		switch k {
		case 0:
			return 0
		case 1:
			return 1
		case 2:
			return 1
		}
	}
	return fibonacci(k-1) + fibonacci(k-2)
}
