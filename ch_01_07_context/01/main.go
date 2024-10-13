package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// выполняет функцию fn
func execute(cancel <-chan struct{}, fn func() int) (int, error) {
	ch := make(chan int, 1)

	go func() {
		ch <- fn()
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-cancel:
		return 0, errors.New("canceled")
	}
}

func main() {
	// работает в течение 100 мс
	work := func() int {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("work done")
		return 42
	}

	// ждет 50 мс, после этого
	// с вероятностью 50% отменяет работу
	maybeCancel := func(cancel chan struct{}) {
		time.Sleep(50 * time.Millisecond)
		if rand.Float32() < 0.5 {
			close(cancel)
		}
	}

	cancel := make(chan struct{})

	go maybeCancel(cancel)

	res, err := execute(cancel, work)
	fmt.Println(res, err)
}
