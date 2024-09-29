package main

import (
	"fmt"
	"strings"
	"unicode"
)

// nextFunc возвращает следующее слово из генератора
type nextFunc9 func() string

// counter хранит количество цифр в каждом слове.
// ключ карты - слово, а значение - количество цифр в слове.
type counter9 map[string]int

// pair хранит слово и количество цифр в нем
type pair9 struct {
	word  string
	count int
}

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWords9(next nextFunc9) counter9 {

	// начало решения

	// отправляет слова на подсчет
	pending := make(chan string)
	go submitWords(next, pending)

	// считает цифры в словах
	counted := make(chan pair9)
	go countWords(pending, counted)

	return fillStats(counted)
	// конец решения
}

func submitWords(next nextFunc9, pending chan<- string) {
	for {
		word := next()
		pending <- word
		if word == "" {
			break
		}
	}
}

func countWords(pending chan string, counted chan pair9) {
	for {
		word := <-pending
		count := countDigits9(word)
		counted <- pair9{word: word, count: count}
		if word == "" {
			break
		}
	}
}

func fillStats(counted chan pair9) counter9 {
	stats := make(counter9)
	for {
		wc := <-counted
		if wc.word == "" {
			break
		}
		stats[wc.word] = wc.count
	}
	return stats
}

// countDigits9 возвращает количество цифр в строке
func countDigits9(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats9 печатает слова и количество цифр в каждом
func printStats9(stats counter9) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// wordGenerator9 возвращает генератор, который выдает слова из фразы
func wordGenerator9(phrase string) nextFunc9 {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator9(phrase)
	stats := countDigitsInWords9(next)
	printStats9(stats)
}
