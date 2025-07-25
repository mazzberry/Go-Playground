package main

import "fmt"

type Runner interface {
	Run()
}

type Walker interface {
	Walk()
}

type Shooter interface {
	Shoot()
}

type Player struct {
	Name     string
	age      uint8
	Height   int
	Weight   int
	Position string
}

func main() {
	player1 := &Player{
		Name:     "Mojtaba",
		age:      20,
		Height:   180,
		Weight:   90,
		Position: "Forward",
	}

	var runner Runner = player1
	var walker Walker = player1
	var Shooter Shooter = player1

	runner.Run()
	walker.Walk()
	Shooter.Shoot()
}

func (player *Player) Run() {
	fmt.Printf("\nname : %s , position: %s player is running\n", player.Name, player.Position)
}

func (player *Player) Walk() {
	fmt.Printf("\nname : %s , position: %s player is Walking\n", player.Name, player.Position)
}

func (player *Player) Shoot() {
	fmt.Printf("\nname : %s , position: %s player Shooting\n", player.Name, player.Position)
}
