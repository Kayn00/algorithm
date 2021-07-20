package main

import (
	"fmt"
	"sync"
)

// 实现的双端列表和 Golang 标准库 container/list 中实现的不一样，感兴趣的可以阅读标准库的实现。
type (
	// 双端列表，双端队列
	DoubleList struct {
		head *DoubleListNode
		tail *DoubleListNode
		len  int
		lock sync.Mutex
	}

	// 列表节点
	DoubleListNode struct {
		pre   *DoubleListNode
		next  *DoubleListNode
		value string
	}
)

// 获取节点值
func (node *DoubleListNode) GetValue() string {
	return node.value
}

// 获取节点前驱节点
func (node *DoubleListNode) GetPre() *DoubleListNode {
	return node.pre
}

// 获取节点后驱节点
func (node *DoubleListNode) GetNext() *DoubleListNode {
	return node.next
}

// 是否存在后驱节点
func (node *DoubleListNode) HashNext() bool {
	return node.next != nil
}

// 是否存在前驱节点
func (node *DoubleListNode) HashPre() bool {
	return node.pre != nil
}

// 是否为空节点
func (node *DoubleListNode) IsNil() bool {
	return node == nil
}

// 返回列表长度
func (list *DoubleList) Len() int {
	return list.len
}

// 返回列表链表头节点
func (list *DoubleList) First() *DoubleListNode {
	return list.head
}

// 返回列表链表尾节点
func (list *DoubleList) Last() *DoubleListNode {
	return list.tail
}

// 添加节点到链表头部的第N个元素之前，N=0表示新节点成为新的头部
func (list *DoubleList) AddNodeFromHead(n int, v string) {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 索引超过列表长度，一定找不到，panic
	if n > list.len {
		panic("index out")
	}

	// 先找出头部
	node := list.head

	// 往后遍历拿到第N+1个位置的元素
	for i := 1; i <= n; i++ {
		node = node.next
	}

	// 新节点
	newNode := new(DoubleListNode)
	newNode.value = v

	// 如果定位到的节点为空，表示列表为空，将新节点设置为新头部和新尾部
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 定位到的节点
		pre := node.pre

		// 如果定位到的节点前驱为nil，那么定位到的节点为链表头部，需要换头部
		if pre.IsNil() {
			// 将新节点链接在老头部之前
			newNode.next = node
			node.pre = newNode
			// 新节点成为头部
			list.head = newNode
		} else {
			// 将新节点插入到定位到的节点之前
			// 定位到的节点的前驱节点pre现在链接到新节点上
			pre.next = newNode
			newNode.pre = pre

			// 定位到的节点的后驱节点node.next 现在链接到新节点上
			node.next.pre = newNode
			newNode.next = node.next
		}
	}

	// 列表长度+1
	list.len = list.len + 1
}

// 添加节点到链表尾部的第N个元素之后，N=0表示新节点成为新的尾部
func (list *DoubleList) AddNodeFromTail(n int, v string) {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 索引超过列表长度，一定找不到，panic
	if n > list.len {
		panic("index out")
	}

	// 先找出尾部
	node := list.tail

	// 往前遍历拿到第N+1个位置的元素
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	// 新节点
	newNode := new(DoubleListNode)
	newNode.value = v

	// 如果定位到的节点为空，表示列表为空，将新节点设置为新头部和新尾部
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 定位到的节点，它的后驱
		next := node.next

		// 如果定位到的节点后驱为nil，那么定位到的节点为链表尾部，需要换尾部
		if next.IsNil() {
			// 将新节点链接在老尾部之后
			node.next = newNode
			newNode.pre = node

			// 新节点成为尾部
			list.tail = newNode
		} else {
			// 将新节点插入到定位到的节点之后
			// 新节点链接到定位的节点之后
			newNode.pre = node
			node.next = newNode

			// 定位到的节点的后驱节点链接在新节点之后
			newNode.next = next
			next.pre = newNode
		}
	}

	// 列表长度+1
	list.len = list.len + 1
}

// 从头部开始往后找，获取第N+1个位置的节点，索引从0开始
func (list *DoubleList) IndexFormHead(n int) *DoubleListNode {
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}

	// 获取头部节点
	node := list.head

	// 往后遍历拿到第N+1个位置的元素
	for i := 1; i <= n; i++ {
		node = node.next
	}

	return node
}

// 从尾部开始某个位置获取列表节点
func (list *DoubleList) IndexFromTail(n int) *DoubleListNode {
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}

	// 获取尾部节点
	node := list.tail

	// 往前遍历拿到第N+1个位置的元素
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	return node
}

// 从头部开始往后找，获取第N+1个位置的节点，并移除返回
func (list *DoubleList) PopFromHead(n int) *DoubleListNode {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}

	// 获取头部
	node := list.head

	// 往后遍历拿到第N+1个位置的元素
	for i := 1; i <= n; i++ {
		node = node.next
	}

	// 移除的节点的前驱和后驱
	pre := node.pre
	next := node.next

	// 如果前驱和后驱都为nil,那么移除的节点为链表唯一节点
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		// 表示移除的是头部节点，那么下一个节点成为头节点
		list.head = next
		next.pre = nil
	} else if next.IsNil() {
		// 表示移除的是尾部节点，那么上一个节点成为尾节点
		list.tail = pre
		pre.next = nil
	} else {
		// 移除的是中间节点
		pre.next = next
		next.pre = pre
	}

	// 节点减一
	list.len = list.len - 1
	return node
}

// 从尾部开始往前找，获取第N+1个位置的节点，并移除返回
func (list *DoubleList) PopFromTail(n int) *DoubleListNode {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}

	// 获取尾部
	node := list.tail

	// 往前遍历拿到第N+1个位置的元素
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	// 移除的节点的前驱和后驱
	pre := node.pre
	next := node.next

	// 如果前驱和后驱都为nil，那么移除的节点为链表唯一节点
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		// 表示移除的是头部节点，那么下一个节点成为头节点
		list.head = next
		next.pre = nil
	} else if next.IsNil() {
		// 表示移除的是尾部节点，那么上一个节点成为尾节点
		list.tail = pre
		pre.next = nil
	} else {
		// 移除的是中间节点
		pre.next = next
		next.pre = pre
	}

	// 节点减一
	list.len = list.len - 1
	return node
}

func main() {
	list := new(DoubleList)
	// 在列表头部插入新元素
	list.AddNodeFromHead(0, "i")
	list.AddNodeFromHead(0, "love")
	list.AddNodeFromHead(0, "you")
	// 在列表尾部插入新元素
	list.AddNodeFromTail(0, "may")
	list.AddNodeFromTail(0, "happy")

	// 正常遍历，比较慢
	for i := 0; i < list.Len(); i++ {
		// 从头部开始索引
		node := list.IndexFormHead(i)

		// 节点为空不可能，因为list.Len()使得索引不会越界
		if !node.IsNil() {
			fmt.Println(node.GetValue())
		}
	}
	fmt.Println("--------------")

	// 正常遍历，特别快
	// 先取出第一个元素
	first := list.First()
	for !first.IsNil() {
		// 如果非空就一直遍历
		fmt.Println(first.GetValue())
		// 接着下一个节点
		first = first.GetNext()
	}
	fmt.Println("--------------")

	// 元素一个个pop出来
	for {
		node := list.PopFromHead(0)
		if node.IsNil() {
			// 没有元素了，直接返回
			break
		}
		fmt.Println(node.GetValue())
	}
	fmt.Println("---------------")
	fmt.Println("len", list.Len())
}
