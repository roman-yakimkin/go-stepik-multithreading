package main

import (
	"fmt"
	"time"
)

func withWorkers(n int, fn func()) (handle func(), wait func()) {
	// канал с токенами
	free := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		free <- struct{}{}
	}

	// выполняет fn, но не более n одновременно
	handle = func() {
		<-free
		go func() {
			fn()
			free <- struct{}{}
		}()
	}

	// ожидает, пока все запущенные fn отработают
	wait = func() {
		for i := 0; i < n; i++ {
			<-free
		}
	}

	return handle, wait
}

func main() {
	work := func() {
		time.Sleep(100 * time.Millisecond)
	}

	handle, wait := withWorkers(2, work)

	start := time.Now()

	handle()
	handle()
	handle()
	handle()
	wait()

	fmt.Println("4 calls took", time.Since(start))
}
