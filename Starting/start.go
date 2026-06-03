package main

import (
	"fmt"
	"time"
)

func Start() {
	var name string

	fmt.Printf("Приветсвую вас в игре!!!\nДайте имя персонажу: ")
	fmt.Scan(&name)
	player := NewPlayer(name)
	unit := NewUnit()

	fmt.Printf("Приветсвую %s\nВаш HP: %d\nВаш урон: %d", player.Name, player.HP, player.Damage)

	time.Sleep(time.Second * 2)
	fmt.Print("\033[H\033[2J")

	fmt.Println("Ну что начнем!!!")
	time.Sleep(time.Second * 2)

	fmt.Print("\033[H\033[2J")

	for {
		fmt.Println("1: Настройки | 2: Левл | 3: Exit")
		var b int16
		fmt.Scan(&b)
		time.Sleep(time.Second * 2)
		fmt.Print("\033[H\033[2J")

		if b == 3 {
			break
		} else if b == 2 {
			Levl_Start(unit, player)
		} else if b == 1 {
			Setting(player)
		}
	}

}