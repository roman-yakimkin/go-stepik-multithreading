// Конкурентно-безопасный стек.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// начало решения

// Stack представляет конкурентно-безопасный стек без блокировок.
type Stack struct {
	top atomic.Pointer[Node]
}

// Node представляет элемент стека.
type Node struct {
	val  int
	next *Node
}

// Push добавляет значение на вершину стека.
func (s *Stack) Push(val int) {
	node := &Node{
		val: val,
	}
	for {
		top := s.top.Load()
		node.next = top
		if s.top.CompareAndSwap(top, node) {
			return
		}
	}
}

// Pop удаляет и возвращает вершину стека.
// Если стек пуст, возвращает false.
func (s *Stack) Pop() (int, bool) {
	for {
		top := s.top.Load()
		if top == nil {
			return 0, false // стек пуст
		}
		if s.top.CompareAndSwap(top, top.next) {
			return top.val, true
		}
	}
}

// конец решения

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)

	stack := &Stack{}
	for i := range 1000 {
		go func() {
			time.Sleep(time.Millisecond)
			stack.Push(i)
			wg.Done()
		}()
	}

	wg.Wait()
	count := 0
	for _, ok := stack.Pop(); ok; _, ok = stack.Pop() {
		count++
	}
	fmt.Println(count)
}
