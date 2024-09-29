package main

import (
	"fmt"
	"strings"
	"unicode"
)

// counter6 количество цифр в каждом слове
type counter6 map[string]int

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWords6(next)
	printStats6(stats)
}

// nextFunc возвращает следующее слово из генератора.
type nextFunc func() string

// countDigits возвращает количество цифр в строке
func countDigits6(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func wordGenerator(phrase string) nextFunc {
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

func countDigitsInWords6(next nextFunc) counter6 {
	stats := make(counter6)
	for {
		word := next()
		if word == "" {
			break
		}
		count := countDigits6(word)
		stats[word] = count
	}
	return stats
}

func printStats6(stats counter6) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}
