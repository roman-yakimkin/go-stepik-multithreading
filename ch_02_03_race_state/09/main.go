// Ограничитель вызовов
package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrBusy = errors.New("busy")
var ErrCanceled = errors.New("canceled")

// начало решения

// throttle следит, чтобы функция fn выполнялась не более limit раз в секунду.
// Возвращает функции handle (выполняет fn с учетом лимита) и cancel (останавливает ограничитель).
func throttle(limit int, fn func()) (handle func() error, cancel func()) {
	var canceled bool
	var mu sync.Mutex
	var counter int
	ticker := time.NewTicker(time.Second)

	done := make(chan struct{})

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case <-done:
					return
				default:
					mu.Lock()
					counter = 0
					mu.Unlock()
				}
			case <-done:
				return
			}
		}
	}()

	handle = func() error {
		mu.Lock()
		defer mu.Unlock()
		if canceled {
			return ErrCanceled
		}

		select {
		case <-done:
			return nil
		default:
			counter++
			if counter > limit {
				return ErrBusy
			}
			fn()
		}

		return nil
	}

	cancel = func() {
		mu.Lock()
		defer mu.Unlock()
		if !canceled {
			canceled = true
			close(done)
		}
	}

	return handle, cancel
}

// конец решения

func main() {
	work := func() {
		fmt.Print(".")
	}

	handle, cancel := throttle(5, work)
	defer cancel()

	const n = 8
	var nOK, nErr int
	for i := 0; i < n; i++ {
		err := handle()
		if err == nil {
			nOK += 1
		} else {
			nErr += 1
		}
	}
	fmt.Println()
	fmt.Printf("%d calls: %d OK, %d busy\n", n, nOK, nErr)
}
