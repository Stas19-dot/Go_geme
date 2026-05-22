package main

import (
	"fmt"
	"time"
)

func Start() {
	var name string

	fmt.Printf("Приветсвую вас в игре!!!\nДайте имя персонажу: ")
	fmt.Scan(&name)
	player := NewPlayer(name, 100, 10)

	fmt.Printf("Приветсвую %s\nВаш HP: %d\nВаш урон: %d", player.Name, player.HP, player.Damage)
	time.Sleep(time.Second * 2)
	fmt.Print("\033[H\033[2J")
	fmt.Println("Ну что начнем!!!")
	

}
