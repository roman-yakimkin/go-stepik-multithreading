// Барьер синхронизации.
package main

import (
	"fmt"
	"sync"
	"time"
)

// начало решения

// Barrier представляет барьер синхронизации.
type Barrier struct {
	n    int
	cnt  int
	cond *sync.Cond
}

// NewBarrier создает новый барьер с указанным порогом.
func NewBarrier(n int) *Barrier {
	return &Barrier{
		n:    n,
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

// Touch фиксирует, что вызывающая горутина достигла барьера.
// Если барьера достигли меньше n горутин, Touch блокирует вызывающую горутину.
// Когда n горутин достигнут барьера, Touch разблокирует их все.
// Если барьер уже разблокирован, Touch не блокирует вызывающую горутину.
func (b *Barrier) Touch() {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	b.cnt++
	for b.cnt < b.n {
		b.cond.Wait()
	}
	b.cond.Broadcast()
}

// конец решения

func main() {
	const nWorkers = 4
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(nWorkers)

	b := NewBarrier(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func(i int) {
			// эмулируем подготовительный шаг...
			dur := time.Duration((i+1)*10) * time.Millisecond
			time.Sleep(dur)
			fmt.Printf("ready to go after %d ms\n", dur.Milliseconds())

			// ждем, пока все горутины соберутся у барьера
			b.Touch()

			// эмулируем основной шаг...
			fmt.Println("go!")
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("all done in %d ms\n", time.Since(start).Milliseconds())

	/*
		ready to go after 10 ms
		ready to go after 20 ms
		ready to go after 30 ms
		ready to go after 40 ms
		go!
		go!
		go!
		go!
		all done in 40 ms
	*/

}
