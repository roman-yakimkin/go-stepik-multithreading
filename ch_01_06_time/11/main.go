// Ограничитель скорости
package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCanceled error = errors.New("canceled")

// начало решения

// throttle следит, чтобы функция fn выполнялась не более limit раз в секунду.
// Возвращает функции handle (выполняет fn с учетом лимита) и cancel (останавливает ограничитель).
func throttle(limit int, fn func()) (handle func() error, cancel func()) {
	dur := time.Duration(int64(time.Second) / int64(limit))
	done := make(chan struct{})
	qu := make(chan struct{})
	var cancelled bool
	ticker := time.NewTicker(dur)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case _, ok := <-qu:
					if ok {
						go fn()
					}
				case <-done:
					return
				}
			case <-done:
				return
			}
		}
	}()
	handle = func() error {
		if cancelled {
			return ErrCanceled
		}
		qu <- struct{}{}
		return nil
	}
	cancel = func() {
		if !cancelled {
			cancelled = true
			close(done)
			close(qu)
		}
	}
	return handle, cancel
}

// конец решения

func main() {
	work := func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}

	handle, cancel := throttle(5, work)
	defer cancel()

	start := time.Now()
	const n = 10
	for i := 0; i < n; i++ {
		handle()
	}
	fmt.Println()
	fmt.Printf("%d queries took %v\n", n, time.Since(start))
}
