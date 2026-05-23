package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)



// Функция для создания игрока
func NewPlayer(name string) *Player {

	return &Player{Name: name, HP: 100, Damage: 10}

}

// Функция для создания моба
func NewUnit() *Unit {

	return &Unit{HP: 20, Damage: 14}

}

// Добавление моба в одну команду
// func AddUnit(s *Squad, u *Unit) error {
// 	if len((*s).Units) == 5 {
// 		return errors.New("В команде уже максимальное кол-во мобов")

// 		// Проверка на HP
// 	} else if u.HP <= 0 {
// 		return errors.New("Вы не можете добавить моба с HP <= 0")
// 	} else {

// 		s.Units = append(s.Units, *u)
// 		return nil

// 	}
// }

// Функция атаки
func Attack(u *Unit, player *Player) error {
	if u.HP > 0 {
		fmt.Println("Вы атакуете...")

		generator := rand.New(rand.NewSource(time.Now().UnixNano()))
		n := generator.Int63()
		time.Sleep(time.Second * 2)

		if n%2 == 0 {
			fmt.Printf("Поздравляю вы попали. HP -%d\n", player.Damage)
			u.HP -= player.Damage
		} else {
			fmt.Println("К сожалению вы промахнулись...")
		}
		time.Sleep(time.Second)
		fmt.Println("Моб атакует...")

		time.Sleep(time.Second * 2)
		generator = rand.New(rand.NewSource(time.Now().UnixNano()))
		n = generator.Int63()

		if n%5 == 0 {
			fmt.Printf("Моб в вас попал: -%dHP\n", u.Damage)
			(*player).HP -= u.Damage
		} else {
			fmt.Println("Моб промахнулся!!!")
		}

		return nil

	} else {

		return errors.New("Моб уже умер...")

	}
}

func Output(player *Player, unit *Unit) {

	fmt.Println("Ваше HP:", (*player).HP)
	fmt.Println("HP моба:", (*unit).HP)

}

func levl(u *Unit, levl int) {

	u.HP *= levl
	u.Damage *= levl

}

func unload_u(u *Unit) {

	u.Damage = 20
	u.HP = 14

}

func unload_p(p *Player) {

	p.Damage = 10
	p.HP = 100

}