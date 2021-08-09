package main

import (
	"container/list"
	"fmt"
)

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}

	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}

func main() {
	obj := Constructor()
	obj.AppendTail(1)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	obj.AppendTail(2)
	obj.AppendTail(3)
	obj.AppendTail(4)
	fmt.Println(obj.DeleteHead())
	obj.AppendTail(5)
	fmt.Println(obj.DeleteHead())
	obj.AppendTail(6)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
}
