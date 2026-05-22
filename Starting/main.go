package main

func main() {
	player := NewPlayer("Stas", 100, 10)
	a := NewUnit("SF", 20, 42)
	
	Attack(a, player)
	Output(player, a)

}
