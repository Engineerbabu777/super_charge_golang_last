package main

import (
	"fmt"
	"time"
)

func main() {
	// Pointer()
	// CalculateValues(5,5)

	waitingch := make(chan int)

	go FetchResource(&waitingch, 1)
	go FetchResource(&waitingch, 2)
	go FetchResource(&waitingch, 3)

	// time.Sleep(time.Second * 5)

	<-waitingch
	<-waitingch
	<-waitingch

}

type Player struct {
	HP int
}

func TakeDamage(player *Player, amount int) {
	player.HP -= amount
	fmt.Println("player is taking damage -> ", player.HP)
}

func Pointer() {
	player := &Player{
		HP: 100,
	}

	TakeDamage(player, 10)
	fmt.Println(player.HP)
}

func CalculateValues(x int, y int) int {
	return x + y
}

func FetchResource(waitingch *chan int, GR int) {
	fmt.Println("Fetching resource ", GR)

	time.Sleep(time.Second * 2)
	fmt.Println("Waiting for 2 secs,  over and out!")
	// return "Hello"
	*waitingch <- GR
}
