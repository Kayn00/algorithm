package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// 二叉排序树又叫二叉搜索树，二叉查找树
func main() {
	fmt.Println(Sort([]int{11, 2, 5, 30, 100, 10, 1, 70}))
}

// 排序
func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	return appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
