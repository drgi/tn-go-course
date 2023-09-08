package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	maxScore = 11
)

type Game struct {
	wg      *sync.WaitGroup
	players [2]*Player
	table   chan string
}

func NewGame(players [2]*Player) *Game {
	return &Game{wg: new(sync.WaitGroup), players: players, table: make(chan string)}
}

func (g *Game) Play() {
	g.wg.Add(2)
	go g.players[0].play(g.table, g.wg, g.stopGame)
	go g.players[1].play(g.table, g.wg, g.stopGame)
	g.table <- "begin"
	g.wg.Wait()
}

func (g *Game) stopGame(winner *Player) {
	close(g.table)
	fmt.Println(
		"Победа игрока: ", winner.name,
		"Со счетом: ", g.players[0].goals, g.players[1].goals)
}

type Player struct {
	name  string
	hod   string
	goals int
}

func (p *Player) incrementGoal() {
	p.goals++
}

func (p *Player) isWinner() bool {
	return p.goals == maxScore
}

func (p *Player) bey(table chan string, hit string) {
	table <- hit
}

func (p *Player) play(table chan string, wg *sync.WaitGroup, stop func(*Player)) {
	defer wg.Done()
	for v := range table {
		var hit string
		// выбор первого удара
		switch v {
		case "begin", "pong", "stop":
			hit = "ping"
		case "ping":
			hit = "pong"
		}
		fmt.Println("Удар от игрока: ", p.name, hit)
		// получает ли игрок очко?
		if point() {
			p.incrementGoal()
			fmt.Printf("Игрок: %s, получает %v очко!\n", p.name, p.goals)
			// было ли очко победным?
			if p.isWinner() {
				stop(p)
				return
			}
			table <- "stop"
		} else {
			// удар игрока который ходит в текущий момент
			p.bey(table, hit)
		}

	}
}

func main() {
	p1 := &Player{name: "Bob", hod: "ping"}
	p2 := &Player{name: "Jon", hod: "pong"}
	players := [2]*Player{p1, p2}

	game := NewGame(players)
	game.Play()

}

func point() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(5) == 3
}
