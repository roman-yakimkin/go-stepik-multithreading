package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

type pair struct {
	src      string
	reversed string
}

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
		defer fmt.Println("generate done...")
		defer close(out)
		for {
			select {
			case <-cancel:
				return
			case out <- randomWord(5):
			}
		}
	}()
	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	isUnique := func(word string) bool {
		words := make(map[rune]struct{})
		cnt := 0
		for _, r := range word {
			words[r] = struct{}{}
			cnt++
		}
		return len(words) == cnt
	}

	go func() {
		defer fmt.Println("take unique done...")
		defer close(out)
		for {
			select {
			case <-cancel:
				return
			case word, ok := <-in:
				if !ok {
					return
				}
				if isUnique(word) {
					select {
					case <-cancel:
						return
					case out <- word:
					}
				}
			}
		}
	}()
	return out
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan pair {
	out := make(chan pair)

	reverseWord := func(word string) pair {
		var reversed string
		for _, r := range word {
			reversed = string(r) + reversed
		}
		return pair{
			src:      word,
			reversed: reversed,
		}
	}
	go func() {
		defer fmt.Println("reverse done ...")
		defer close(out)

		for word := range in {
			select {
			case <-cancel:
				return
			case out <- reverseWord(word):
			}
		}
	}()
	return out
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan pair) <-chan pair {
	out := make(chan pair)
	go func() {
		defer fmt.Println("merge done...")
		defer close(out)

		for c1 != nil || c2 != nil {
			select {
			case <-cancel:
				return
			case val1, ok := <-c1:
				if ok {
					select {
					case out <- val1:
					case <-cancel:
						return
					}
				} else {
					c1 = nil
				}
			case val2, ok := <-c2:
				if ok {
					select {
					case out <- val2:
					case <-cancel:
						return
					}
				} else {
					c2 = nil
				}
			}
		}
	}()
	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan pair, n int) {
	defer fmt.Println("print done...")
	cnt := 0
	for val := range in {
		select {
		case <-cancel:
			return
		default:
			if cnt < n {
				fmt.Printf("%s -> %s\n", val.src, val.reversed)
				cnt++
			} else {
				return
			}
		}
	}
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)

	print(cancel, c4, 10)
	time.Sleep(time.Millisecond * 1000)
}
