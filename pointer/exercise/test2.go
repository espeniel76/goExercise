package exercise

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	p := &Actor{Name: name, HP: hp, Speed: speed}
	// p := &Actor{name, hp, speed}
	return p
}

func Exercise2() {
	var actor = NewActor("금토끼", 99, 100)
	fmt.Println(actor.Name)
	fmt.Println(actor.HP)
	fmt.Println(actor.Speed)
}
