package main

import "fmt"

type sortTree struct {
	value       int
	left, right *sortTree
}

func main() {
	//var t *sortTree
	//arr := []int{62, 68, 58, 47, 35, 73, 51, 99, 37, 93}
	//for i := 0; i < len(arr); i++ {
	//	insertBst(t, arr[i])
	//}
	//fmt.Println(t)
	//
	//fmt.Println(sortBst([]int{}, t))
	var t *sortTree
	insertBst(t, 62)
	fmt.Println(t)
}

// 排序
func sortBst(values []int, t *sortTree) []int {
	if t != nil {
		values = sortBst(values, t.left)
		values = append(values, t.value)
		values = sortBst(values, t.right)
	}
	return values
}

// 递归查找二叉排序树t中是否存在key
// 指针f指向t的父节点，其初始值为nil
// 若查找成功，则指针p指向该数据元素结点，并返回true
// 否则指针p指向查找路径上访问的最后一个结点并返回false
func searchBst(t, f, p *sortTree, key int) bool {
	if t == nil {
		p = f
		return false
	} else if key == t.value {
		p = t
		return true
	} else if key < t.value {
		return searchBst(t.left, t, p, key)
	} else {
		return searchBst(t.right, t, p, key)
	}
}

// 当二叉排序树t中不存在关键字等于key的数据元素时
// 插入key并返回true，否则返回false
func insertBst(t *sortTree, key int) bool {
	var p *sortTree
	s := new(sortTree)
	if !searchBst(t, nil, p, key) {
		fmt.Println("11111")
		s.value = key
		if p == nil {
			fmt.Println("2222")
			fmt.Println("s is : ", s)
			t = s
		} else if key < p.value {
			p.left = s
		} else {
			p.right = s
		}
		return true
	} else {
		return false
	}
}
