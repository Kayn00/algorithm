package main

import "fmt"

func InsertSort(list []int) {
	n := len(list)
	// 进行N-1轮迭代
	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置

		// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
		if deal < list[j] {
			// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j] // 某数后移，给待排序留空位
			}
			list[j+1] = deal // 结束了，待排序的数插入空位
		}
	}
}

// 自底向上归并排序优化版本
func MergeSort3(array []int, n int) {
	// 按照三个元素为一组进行小数组排序，使用直接插入排序
	blockSize := 3
	a, b := 0, blockSize
	for b <= n {
		InsertSort(array[a:b])
		a = b
		b += blockSize
	}
	InsertSort(array[a:n])

	// 将这些小数组进行归并
	for blockSize < n {
		a, b = 0, 2*blockSize
		for b <= n {
			merge3(array, a, a+blockSize, b)
			a = b
			b += 2 * blockSize
		}
		if m := a + blockSize; m < n {
			merge3(array, a, m, n)
		}
		blockSize *= 2
	}
}

// 原地归并操作
func merge3(array []int, begin, mid, end int) {
	// 三个下标，将数组array[begin,mid]和array[mid,end-1]
	i, j, k := begin, mid, end-1 // 因为数组下标从0开始，所以k=end-1

	for j-i > 0 && k-j >= 0 {
		step := 0
		// 从i向右移动，找到第一个array[i]>array[j]的索引
		for j-i > 0 && array[i] <= array[j] {
			i++
		}

		// 从j向右移动，找到第一个array[j]>array[i]的索引
		for k-j >= 0 && array[j] <= array[i] {
			j++
			step++
		}

		fmt.Println("i:", i, " j:", j, " step:", step, " k:", k)
		// 进行手摇翻转，将array[i,mid]和[mid,j-1]进行位置互换
		// mid 是从j开始向右出发的，所以mid=j-step
		rotation(array, i, j-step, j-1)
		fmt.Println(array)
	}
}

// 手摇算法，将array[l,l+1,l+2,...,mid-2,mid-1,mid,mid+1,mid+2,...,r-2,r-1,r]从mid开始两边交换位置
// 1.先逆序前部分:array[mid-1,mid-2,...,l+2,l+1,l]
// 2.后逆序后部分:array[r,r-1,r-2,...mid+2,mid+1,mid]
// 3.上两步完成后:array[mid-1,mid-2,...,l+2,l+1,l,r,r-1,r-2,...,mid+2,mid+1,mid]
// 4.整体逆序:array[mid,mid+1,mid+2,...,r-2,r-1,r,l+1,l+2,...,mid-2,mid-1]
func rotation(array []int, l, mid, r int) {
	reverse(array, l, mid-1)
	reverse(array, mid, r)
	reverse(array, l, r)
}

func reverse(array []int, l, r int) {
	for l < r {
		// 左右相互交换
		array[l], array[r] = array[r], array[l]
		l++
		r--
	}
}

func main() {
	//list := []int{5}
	//MergeSort3(list, len(list))
	//fmt.Println(list)
	//
	//list1 := []int{5, 9}
	//MergeSort3(list1, len(list1))
	//fmt.Println(list1)

	//list2 := []int{3, 6, 9, 5, 7, 8}
	//MergeSort3(list2, len(list2))
	//fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MergeSort3(list3, len(list3))
	//fmt.Println(list3)
	//
	//list4 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 45, 67, 2, 5, 24, 56, 34, 24, 56, 2, 2, 21, 4, 1, 4, 7, 9}
	//MergeSort3(list4, len(list4))
	//fmt.Println(list4)
}
