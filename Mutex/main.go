package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	mu sync.RWMutex
	health int
}

func NewPlayer() *Player{
	return &Player{
		health: 100,
	}
}

func startUILoop(p *Player){
	ticker := time.NewTicker(time.Second)
	for {
		// reading from state
		p.mu.RLock()
		fmt.Printf("Player health: %d\r", p.health)
		p.mu.RLock()
		<-ticker.C // wait ticker to complete
	}
}

func startGameLoop(p *Player){
	ticker := time.NewTicker(time.Millisecond * 3000)
	for {
		p.mu.Lock()
		p.health -= rand.Intn(40) // rand number betwen 0 and 40

		if p.health <= 0 {
			fmt.Print("GAme over")
			break
		}
		p.mu.Unlock()
		<-ticker.C
	}
}


func main(){
	player := NewPlayer()
	go startUILoop(player)
	startGameLoop(player)
}