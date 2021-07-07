package main

import "fmt"

type Ring struct {
	next, prev *Ring
	Value      interface{}
}

// 初始化空的循环链表，前驱和后驱都指向自己
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// 创建n个节点的循环链表
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}

	r := new(Ring)
	r.Value = 0
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
		p.Value = i
	}
	p.next = r
	r.prev = p
	return r
}

// 获取下一个节点
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// 获取第n个节点 因为链表是循环的，当n为负数，表示从前面往前遍历，否则往后面遍历
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// 往节点r，链接一个节点，并且返回之前节点r的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// 删除节点后面的n个节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	//fmt.Println(r.Move(n + 1).Value)
	return r.Link(r.Move(n + 1))
}

func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}

func main() {
	//r := New(5)
	//	//fmt.Println(r.Value)
	//	//next := r.Next()
	//	//fmt.Println(next.Value)
	//	//move := r.Move(3)
	//	//fmt.Println(move.Value)
	//linkNewTest()
	deleteTest()
}

func linkNewTest() {
	// 第一个节点
	r := &Ring{Value: -1}

	//
	r1 := &Ring{Value: 2}
	r1.Link(&Ring{Value: 1})

	// 添加新的五个节点
	r.Link(r1)
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	r.Link(&Ring{Value: 5})

	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)

		// 移到下一个节点
		node = node.Next()

		// 如果节点回到了起点，结束
		if node == r {
			return
		}
	}
}

func deleteTest() {
	// 第一个节点
	r := &Ring{Value: -1}

	//
	r1 := &Ring{Value: 2}
	r1.Link(&Ring{Value: 1})

	// 添加新的五个节点
	r.Link(r1)
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	r.Link(&Ring{Value: 5})

	fmt.Println("r len is : ", r.Len())

	temp := r.Unlink(3)

	node := r
	for i := 0; i < 10; i++ {
		// 打印节点值
		fmt.Println(node.Value)

		// 移到下一个节点
		node = node.Next()

		// 如果节点回到了起点，结束
		if node == r {
			break
		}
	}

	fmt.Println("---------")

	node = temp
	for i := 0; i <= 10; i++ {
		// 打印节点值
		fmt.Println(node.Value)

		// 移到下一个节点
		node = node.Next()

		// 如果节点回到了起点，结束
		if node == temp {
			break
		}
	}
}
