package main

import (
	"fmt"
	"time"
)

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// начало решения

	type pair struct {
		index int
		value any
	}

	result := make([]any, len(funcs))
	values := make(chan pair, len(funcs))

	// выполните все переданные функции,
	// соберите результаты в срез
	// и верните его
	for i, fn := range funcs {
		go func(idx int) {
			values <- pair{
				index: idx,
				value: fn(),
			}
		}(i)
	}

	for i := 0; i < len(funcs); i++ {
		pair := <-values
		result[pair.index] = pair.value
	}

	// конец решения
	return result
}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
