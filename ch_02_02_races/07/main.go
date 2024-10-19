// Безопасная карта
package main

import (
	"fmt"
	"sync"
)

// начало решения

// Counter представляет безопасную карту частот слов.
// Ключ - строка, значение - целое число.
type Counter struct {
	co map[string]int
	mu sync.Mutex
}

// Increment увеличивает значение по ключу на 1.
func (c *Counter) Increment(str string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.co[str] = c.co[str] + 1
}

// Value возвращает значение по ключу.
func (c *Counter) Value(str string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.co[str]
}

// Range проходит по всем записям карты,
// и для каждой вызывает функцию fn, передавая в нее ключ и значение.
func (c *Counter) Range(fn func(key string, val int)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.co {
		fn(k, v)
	}
}

// NewCounter создает новую карту частот.
func NewCounter() *Counter {
	return &Counter{
		co: make(map[string]int),
		mu: sync.Mutex{},
	}
}

// конец решения

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(3)

	increment := func(key string, val int) {
		defer wg.Done()
		for ; val > 0; val-- {
			counter.Increment(key)
		}
	}

	go increment("one", 100)
	go increment("two", 200)
	go increment("three", 300)

	wg.Wait()

	fmt.Println("two:", counter.Value("two"))

	fmt.Print("{ ")
	counter.Range(func(key string, val int) {
		fmt.Printf("%s:%d ", key, val)
	})
	fmt.Println("}")
}
