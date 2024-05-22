package main

import "fmt"

func main() {
	// Pointer()
	// CalculateValues(5,5)
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