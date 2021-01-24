package main

import "fmt"

type properties interface {
	setName(name string)
	getName() string
	setFeet(feet int)
	getFeet() int
	setHasPaws(hasPaws bool)
	getHasPaws() bool
}

type animal struct {
	name    string
	feet    int
	hasPaws bool
}

// Monster is a special form of animals.
type Monster struct {
	animal
	abilities []string
}

func (animal animal) setName(name string) {
	animal.name = name
}

func (animal animal) getName() string {
	return animal.name
}

func (animal animal) setFeet(feet int) {
	animal.feet = feet
}

func (animal animal) getFeet() int {
	return animal.feet
}

func (animal animal) setHasPaws(hasPaws bool) {
	animal.hasPaws = hasPaws
}

func (animal animal) getHasPaws() bool {
	return animal.hasPaws
}

func (animal animal) print() {
	fmt.Println(animal.getName(), "has", animal.getFeet(), "feet.")

	if animal.hasPaws {
		fmt.Println(animal.getName(), "has paws.")
	}
}

func (monster Monster) print() {
	fmt.Println(monster.getName(), "is a monster!!!")
	fmt.Println(monster.getName(), "has:")

	for _, ability := range monster.abilities {
		fmt.Println("-", ability)
	}
}

// Init initializes a monster.
func Init(name string, feet int, hasPaws bool, abilities []string) *Monster {
	return &Monster{
		animal{name, feet, hasPaws},
		abilities,
	}
}

func main() {
	var dragon properties = Init("Dragon", 4, true, []string{"Nuclear blast", "Ice smoke"})

	fmt.Println(dragon.getName())
}
