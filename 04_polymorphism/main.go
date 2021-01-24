package main

import "fmt"

type animal struct {
	name    string
	feet    int
	hasPaws bool
}

type monster struct {
	animal
	abilities []string
}

func (animal animal) print() {
	fmt.Println(animal.name, "has", animal.feet, "feet.")

	if animal.hasPaws {
		fmt.Println(animal.name, "has paws.")
	}
}

func (monster monster) print() {
	fmt.Println(monster.name, "is a monster!!!")
	fmt.Println(monster.name, "has:")

	for _, ability := range monster.abilities {
		fmt.Println("-", ability)
	}
}

func main() {
	dragon := monster{
		animal{"Dragon", 4, true},
		[]string{"Nuclear blast", "Ice smoke"},
	}

	dragon.print()
}
