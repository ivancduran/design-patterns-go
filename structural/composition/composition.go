package composition

import "fmt"

// COMPOSITION WITH STRUCTS
type Athlete struct {
	Name string
	ID   int
}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

func Swim() {
	fmt.Println("Swiming!")
}

type CompositeSwimerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

// --------------------

// COMPOSITION WITH STRUCTS 2
