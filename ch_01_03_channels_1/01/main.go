package main

import (
	"fmt"
	"strings"
)

func main() {
	const eof = "__EOF__"

	str := "one,two,,four"

	in := make(chan string)
	go func() {
		words := strings.Split(str, ",")
		for _, word := range words {
			in <- word
		}
		in <- eof // (1)
	}()

	for {
		word := <-in
		if word == eof { // (2)
			break
		}
		if word != "" {
			fmt.Printf("%s ", word)
		}
	}
}
