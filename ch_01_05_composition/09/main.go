package main

import "fmt"

type Total struct {
	count  int
	amount int
}

// генерит числа в указанном диапазоне
func rangeGen(start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			out <- i
		}
	}()
	return out
}

// выбирает счастливые числа
func takeLucky(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			if num%7 == 0 && num%13 != 0 {
				out <- num
			}
		}
	}()
	return out
}

// объединяет независимые каналы
func merge(channels []<-chan int) <-chan int {
	out := make(chan int)
	// ...
	return out
}

// суммирует числа
func sum(in <-chan int) <-chan Total {
	out := make(chan Total)
	go func() {
		defer close(out)
		total := Total{}
		for num := range in {
			total.amount += num
			total.count++
		}
		out <- total
	}()
	return out
}

// печатает результат
func printTotal(in <-chan Total) {
	total := <-in
	fmt.Printf("Total of %d lucky numbers = %d\n", total.count, total.amount)
}

func main() {
	readerChan := rangeGen(1, 1000)
	luckyChans := make([]<-chan int, 4)
	for i := 0; i < 4; i++ {
		luckyChans[i] = takeLucky(readerChan)
	}
	mergedChan := merge(luckyChans)
	totalChan := sum(mergedChan)
	printTotal(totalChan)
	// Total of 132 lucky numbers = 66066
}
