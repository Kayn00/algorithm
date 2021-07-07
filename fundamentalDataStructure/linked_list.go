package main

import "fmt"

type LinkNode struct {
	Data int64
	Next *LinkNode
}

func main() {
	// 新的节点
	node := new(LinkNode)
	node.Data = 2

	// 新的节点
	node1 := new(LinkNode)
	node1.Data = 3
	node.Next = node1

	// 新的节点
	node2 := new(LinkNode)
	node2.Data = 4
	node1.Next = node2

	// 按顺序打印数据
	nowNode := node
	for {
		if nowNode != nil {
			// 打印节点值
			fmt.Println(nowNode.Data)
			// 获取下一个节点
			nowNode = nowNode.Next
			continue
		}
		break
	}
}
