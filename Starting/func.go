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

// Функция для выбора левла и основной игры
func Levl_Start(unit *Unit, player *Player) {
	fmt.Println("Введите номер лвл (1-10): ")
	var c int
	fmt.Scan(&c)
	if c <= 0 || c > 10 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Ошибка!!!")
		time.Sleep(time.Second)
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Print("\033[H\033[2J")
		levl_Unit(unit, c)
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
					player.bank += count * 10
					break
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
}

func Output(player *Player, unit *Unit) {

	fmt.Println("Ваше HP:", (*player).HP)
	fmt.Println("HP моба:", (*unit).HP)

}

func levl_Unit(u *Unit, levl int) {

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
