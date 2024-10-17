package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}, 5)

	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(50 * time.Millisecond)
			fmt.Print(".")
			done <- struct{}{}
		}()
	}

	for i := 0; i < 5; i++ {
		<-done
	}

	fmt.Println("all done")
}
