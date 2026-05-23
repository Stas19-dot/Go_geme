package main

import "fmt"

func New_damage(pl *Player) {
	var baks int
	fmt.Scan(&baks)
	pl.Damage  *= baks
}

func Setting(pl *Player) {

	for {
		var count int
		fmt.Println("1: Улучшение оружия | 2: Информация")
		fmt.Scan(&count)
		if count == 1 {
			New_damage(pl)
		}
	}

}