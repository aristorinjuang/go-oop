package main

import "fmt"

type animal struct {
	name    string
	feet    int
	hasPaws bool
}

func (animal animal) print() {
	fmt.Println(animal.name, "has", animal.feet, "feet.")

	if animal.hasPaws {
		fmt.Println(animal.name, "has paws.")
	}
}

func main() {
	cat := animal{"Cat", 4, false}

	cat.print()

	var dog = new(animal)

	dog.name = "Dog"
	dog.feet = 4
	dog.hasPaws = true

	dog.print()
}
