// Блокирующая очередь.
package main

import (
	"fmt"
	"sync"
)

// начало решения

// Queue - блокирующая FIFO-очередь.
type Queue struct {
	// TODO: переделать на срез и sync.Cond.
	cond  *sync.Cond
	items []int
}

// NewQueue создает новую очередь.
func NewQueue() *Queue {
	// TODO: очередь должна быть безразмерной.
	return &Queue{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

// Put добавляет элемент в очередь.
// Поскольку очередь безразмерная, никогда не блокируется.
func (q *Queue) Put(item int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.items = append(q.items, item)
	q.cond.Signal()
}

// Get извлекает элемент из очереди.
// Если очередь пуста, блокируется до момента,
// пока в очереди не появится элемент.
func (q *Queue) Get() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.items) == 0 {
		q.cond.Wait()
	}

	val := q.items[0]
	q.items = q.items[1:]
	return val
}

// Len возвращает количество элементов в очереди.
func (q *Queue) Len() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return len(q.items)
}

// конец решения

func main() {
	var wg sync.WaitGroup
	q := NewQueue()

	wg.Add(1)
	go func() {
		for i := range 100 {
			q.Put(i)
		}
		wg.Done()
	}()
	wg.Wait()

	total := 0

	wg.Add(1)
	go func() {
		for range 100 {
			total += q.Get()
		}
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("Put x100, Get x100, Total:", total)
}
