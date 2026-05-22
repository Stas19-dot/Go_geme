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
		fmt.Println("1: Информация | 2: Левл | 3: Exit")
		var b int16
		fmt.Scan(&b)
		time.Sleep(time.Second * 2)
		fmt.Print("\033[H\033[2J")

		if b == 3 {
			break
		} else if b == 2 {
			fmt.Println("Введите номер лвл (1-10): ")
			var c int
			fmt.Scan(&c)
			if c <= 0 {
				fmt.Println("Ошибка!!!")
				continue
			} else {
				levl(unit, c)
				var count int
				for {
					fmt.Println("1: Attack | 2: Exite")
					fmt.Scan(&count)
					fmt.Print("\033[H\033[2J")

					if count == 1 {
						Attack(unit, player)
						if unit.HP <= 0 {
							fmt.Print("\033[H\033[2J")
							fmt.Printf("Вы прошли %d лвл\n", c)
							continue
						}
						if player.HP <= 0 {
							fmt.Println("Вы умерли весь ваш прогресс был снят")
							unload_p(player)
						}
						continue
					} else {
						break
					}
				}
			}
		} else if b == 1 {
			fmt.Printf("Имя: %s; HP: %d; Урон: %d\n", (*player).Name, (*player).HP, (*player).Damage)
			continue
		}
	}

}
