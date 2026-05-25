package main

import "fmt"

func New_damage(pl *Player) {
	var baks int
	fmt.Scan(&baks)
	pl.Damage *= baks
}

func Setting(pl *Player) {

	for {
		var count int
		fmt.Println("1: Улучшение оружия | 2: Информация | 3: Exite")
		fmt.Scan(&count)
		if count == 1 {

			New_damage(pl)

		} else if count == 2 {
			fmt.Print("\033[H\033[2J")
			fmt.Printf("Name: %s; HP: %d; Damage: %d\n", pl.Name, pl.HP, pl.Damage)
		} else if count == 3 {
			fmt.Print("\033[H\033[2J")
			break
		}
	}

}
