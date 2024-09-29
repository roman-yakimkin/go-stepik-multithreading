package main

import (
	"fmt"
	"strings"
	"unicode"
)

// nextFunc возвращает следующее слово из генератора
type nextFunc8 func() string

// counter хранит количество цифр в каждом слове.
// ключ карты - слово, а значение - количество цифр в слове.
type counter8 map[string]int

// pair хранит слово и количество цифр в нем
type pair8 struct {
	word  string
	count int
}

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWords8(next nextFunc8) counter8 {
	pending := make(chan string)
	counted := make(chan pair8)

	// начало решения
	stats := make(counter8)

	// отправляет слова на подсчет
	go func() {
		// Пройдите по словам и отправьте их
		// в канал pending
		for {
			word := next()
			pending <- word
			if word == "" {
				break
			}
		}
	}()

	// считает цифры в словах
	go func() {
		// Считайте слова из канала pending,
		// посчитайте количество цифр в каждом,
		// и запишите его в канал counted
		for {
			word := <-pending
			count := countDigits8(word)
			counted <- pair8{word: word, count: count}
			if word == "" {
				break
			}
		}
	}()

	// Считайте значения из канала counted
	// и заполните stats.
	for {
		wc := <-counted
		if wc.word == "" {
			break
		}
		stats[wc.word] = wc.count
	}

	// В результате stats должна содержать слова
	// и количество цифр в каждом.

	// конец решения

	return stats
}

// countDigits8 возвращает количество цифр в строке
func countDigits8(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats8 печатает слова и количество цифр в каждом
func printStats8(stats counter8) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// wordGenerator8 возвращает генератор, который выдает слова из фразы
func wordGenerator8(phrase string) nextFunc8 {
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
	next := wordGenerator8(phrase)
	stats := countDigitsInWords8(next)
	printStats8(stats)
}
