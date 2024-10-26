// Игра «угадай среднее».
package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// начало решения

// Game представляет игру.
type Game struct {
	stakes chan stake
	win    stake
	once   sync.Once
}

// NewGame создает новую игру на nPlayers игроков.
func NewGame(nPlayers int) *Game {
	return &Game{
		stakes: make(chan stake, nPlayers),
	}
}

// Play принимает ставку от игрока.
func (g *Game) Play(player string, num float64) {
	select {
	case g.stakes <- stake{player, num}:
	default:
	}
}

// Finish завершает игру и определяет победителя.
func (g *Game) Finish() stake {
	g.once.Do(func() {
		g.win = g.decideWinner()
	})

	return g.win
}

// конец решения

// stake представляет ставку игрока.
type stake struct {
	player string
	num    float64
}

// decideWinner определяет победителя игры.
// Победитель - игрок, чья ставка ближе всего к средней.
func (g *Game) decideWinner() stake {
	// собираем все ставки
	var s []stake
	for range len(g.stakes) {
		s = append(s, <-g.stakes)
	}

	// находим среднюю ставку
	total := 0.0
	for _, stake := range s {
		total += stake.num
	}
	avg := total / float64(len(s))

	// побеждает тот, чья ставка ближе всего к средней
	var winner stake
	minDist := math.Inf(1)
	for _, stake := range s {
		if dist := math.Abs(stake.num - avg); dist < minDist {
			minDist = dist
			winner = stake
		}
	}

	return winner
}

func main() {
	// создаем новую игру
	game := NewGame(3)

	// игроки делают ставки
	go game.Play("Alice", 10)
	go game.Play("Bob", 21)
	go game.Play("Cindy", 30)
	time.Sleep(10 * time.Millisecond)

	// завершаем игру
	go game.Finish()
	go game.Finish()
	time.Sleep(10 * time.Millisecond)
	winner := game.Finish()

	// оглашаем победителя
	time.Sleep(10 * time.Millisecond)
	fmt.Println("winner:", winner)
	// winner: {Bob 21}
}
