package sort

import "fmt"

// 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// 针对所有的元素重复以上的步骤，除了最后一个。
// 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

func Bubble() {
	numbers := []int{5, 8, 2, 7, 11, 15, 19, 17, 12, 14, 29, 35, 6, 20}
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	fmt.Println(numbers)
}
