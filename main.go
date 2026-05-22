package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Создание структур
type Unit struct {
	Name   string
	HP     int
	Damage int
}

type Squad struct {
	Units       []Unit
	unit_active int
}

type Player struct {
	Name   string
	HP     int
	Damage int
}

//

// Функция создающая нового моба
func NewUnit(name string, hp int, damage int) *Unit {

	return &Unit{Name: name, HP: hp, Damage: damage}

}

// Добавление моба в одну команду
func AddUnit(s *Squad, u *Unit) error {
	if len((*s).Units) == 5 {
		return errors.New("В команде уже максимальное кол-во мобов")

		// Проверка на HP
	} else if u.HP <= 0 {
		return errors.New("Вы не можете добавить моба с HP <= 0")
	} else {

		s.Units = append(s.Units, *u)
		return nil

	}
}

func Attack(s *Squad, num_pleer int, damage int, player *Player) {
	if s.Units[num_pleer].HP > 0 {
		fmt.Println("Вы атакуете...")

		generator := rand.New(rand.NewSource(time.Now().UnixNano()))
		n := generator.Int63()
		time.Sleep(time.Second * 2)

		if n%2 == 0 {
			fmt.Printf("Поздравляюв вы попали. HP -%d\n", player.Damage)
			s.Units[num_pleer].Damage -= player.Damage
		}

		fmt.Println("Моб атакует...")

		time.Sleep(time.Second * 2)
		generator = rand.New(rand.NewSource(time.Now().UnixNano()))
		n = generator.Int63()

		if n%5 == 0 {
			fmt.Printf("Моб в вас попал: -%dHP\n", s.Units[num_pleer].Damage)

		} else {
			fmt.Println("Моб промахнулся!!!")
		}

	} else {

		fmt.Println("Данный моб уже умер (HP <= 0)")

	}
}

func main() {
	var s Squad
	a := NewUnit("SF", 20, 42)
	AddUnit(&s, a)
	Attack(&s, 0, 30)
	fmt.Println(s.Units[0])

}
