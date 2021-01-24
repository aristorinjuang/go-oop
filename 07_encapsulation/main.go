package main

import "fmt"

type properties interface {
	setName(name string)
	getName() string
	setFeet(feet int)
	getFeet() int
	setHasPaws(hasPaws bool)
	getHasPaws() bool
	Print()
}

// Animal has a name and feet. Maybe has paws too.
type Animal struct {
	name    string
	feet    int
	hasPaws bool
}

// Monster is a special form of animals.
type Monster struct {
	Animal
	abilities []string
}

func (animal Animal) setName(name string) {
	animal.name = name
}

func (animal Animal) getName() string {
	return animal.name
}

func (animal Animal) setFeet(feet int) {
	animal.feet = feet
}

func (animal Animal) getFeet() int {
	return animal.feet
}

func (animal Animal) setHasPaws(hasPaws bool) {
	animal.hasPaws = hasPaws
}

func (animal Animal) getHasPaws() bool {
	return animal.hasPaws
}

// Print all properties of an animal.
func (animal Animal) Print() {
	fmt.Println(animal.getName(), "has", animal.getFeet(), "feet.")

	if animal.hasPaws {
		fmt.Println(animal.getName(), "has paws.")
	}
}

// Print all abilities of a monster.
func (monster Monster) Print() {
	fmt.Println(monster.getName(), "is a monster!!!")
	fmt.Println(monster.getName(), "has:")

	for _, ability := range monster.abilities {
		fmt.Println("-", ability)
	}
}

// Init initializes a monster.
func Init(name string, feet int, hasPaws bool, abilities []string) *Monster {
	return &Monster{
		Animal{name, feet, hasPaws},
		abilities,
	}
}

func main() {
	var dragon properties = Init("Dragon", 4, true, []string{"Nuclear blast", "Ice smoke"})

	dragon.Print()
}
