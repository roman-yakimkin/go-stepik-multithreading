package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("B: sending message...")
		messages <- "ping"
		fmt.Println("B: message sent!")
	}()

	fmt.Println("A: Doing some work...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("A: Ready to receive a message...")

	<-messages

	fmt.Println("A: message received!")
	time.Sleep(100 * time.Millisecond)
}
