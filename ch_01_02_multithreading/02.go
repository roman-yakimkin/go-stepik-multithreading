package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go sayWG(&wg, 1, "go is awesome")

	wg.Add(1)
	go sayWG(&wg, 2, "cats are cute")

	wg.Wait()
}

func sayWG(wg *sync.WaitGroup, id int, phrase string) {
	defer wg.Done()
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}
