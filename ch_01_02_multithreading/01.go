package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	go say(1, "go is awesome")
	go say(2, "cats are cute")
	time.Sleep(500 * time.Millisecond)
}

func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}
