package main

import "fmt"

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
	fmt.Println(player.HP);
}
