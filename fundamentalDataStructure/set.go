package main

import (
	"fmt"
	"sync"
)

type (
	Set struct {
		m   map[int]struct{}
		len int
		sync.RWMutex
	}
)

// 新建一个空集合
// 空结构体不占用内存空间
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m: temp,
	}
}

// 添加一个元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{}
	s.len = len(s.m)
}

// 删除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	s.Unlock()

	// 集合没元素直接返回
	if s.len == 0 {
		return
	}

	delete(s.m, item)
}

// 查看元素是否在集合中
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// 查看集合大小
func (s *Set) Len() int {
	return s.len
}

// 集合是否为空
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// 清除集合所有元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{}
	s.len = 0
}

// 将集合转化为列表
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, s.len)
	for item := range s.m {
		fmt.Println("item is : ", item)
		list = append(list, item)
	}
	return list
}

func main() {
	// 初始化一个容量为5的不可重复集合
	s := NewSet(5)

	s.Add(1)
	s.Add(1)
	s.Add(2)
	fmt.Println("list of all items", s.List())

	s.Clear()
	if s.IsEmpty() {
		fmt.Println("empty")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		fmt.Println("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items", s.List())
}
