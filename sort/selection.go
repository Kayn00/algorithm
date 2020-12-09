package sort

import "fmt"

// 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
// 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
// 重复第二步，直到所有元素均排序完毕。

func Selection() {
	numbers := []int{5, 8, 2, 7, 11, 15, 19, 17, 12, 14, 29, 35, 6, 20}
	for i := 0; i < len(numbers); i++ {
		min := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[min] > numbers[j] {
				min = j
			}
		}
		numbers[i], numbers[min] = numbers[min], numbers[i]
	}
	fmt.Println(numbers)
}
