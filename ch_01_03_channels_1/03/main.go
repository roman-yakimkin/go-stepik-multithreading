package main

import "fmt"

func main() {
	words := []string{"1", "2", "3"}
	for idx, val := range words {
		fmt.Println(idx, val)
	}

	in := make(chan string)
	go func() {
		in <- "1"
		in <- "2"
		in <- "3"
		close(in)
	}()
	for val := range in {
		fmt.Println(val)
	}
}
