package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxScore = 11
)

type Game struct {
	currentPlayerIndex int
	current            *Player
	players            [2]*Player
	table              chan string
}

func NewGame(players [2]*Player) *Game {
	return &Game{players: players, table: make(chan string)}
}

func (g *Game) Play() {
	for v := range g.table {
		// начало игры выбор рандомного игрока
		if v == "start" {
			g.currentPlayerIndex = g.randomPlayerIndex()
		}
		currentPlayer := g.players[g.currentPlayerIndex]

		// удар игрока который ходит в текущий момент
		go func() {
			currentPlayer.bey(g.table)

			// получает ли игрок очко?
			if point() {
				currentPlayer.IncrementGoal()
				fmt.Println("Goal from: ", currentPlayer.name, currentPlayer.goals)
			}

			// было ли очко победным?
			if currentPlayer.IsWinner() {
				g.stopGame()
				return
			}
		}()
		g.switchPlayer()
	}
}

func (g *Game) randomPlayerIndex() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1)
}

func (g *Game) stopGame() {
	close(g.table)
	fmt.Println(
		"Победа игрока: ", g.players[g.currentPlayerIndex].name,
		"Со счетом: ", g.players[0].goals, g.players[1].goals)
}

func (g *Game) switchPlayer() {
	if g.currentPlayerIndex == 0 {
		g.currentPlayerIndex = 1
	} else if g.currentPlayerIndex == 1 {
		g.currentPlayerIndex = 0
	}
}

type Player struct {
	name  string
	hod   string
	goals int
}

func (p *Player) IncrementGoal() {
	p.goals++
}

func (p *Player) IsWinner() bool {
	return p.goals == maxScore
}

func (p *Player) bey(table chan string) {
	fmt.Println("Удар от игрока: ", p.name, p.hod)
	table <- p.hod
}

func main() {
	p1 := &Player{name: "Bob", hod: "ping"}
	p2 := &Player{name: "Jon", hod: "pong"}
	players := [2]*Player{p1, p2}

	game := NewGame(players)
	go func() {
		game.table <- "start"
	}()
	game.Play()

}

func point() bool {
	return rand.Intn(5) == 4
}
