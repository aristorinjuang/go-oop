package main

import "fmt"

type animal struct {
	name    string
	feet    int
	hasPaws bool
}

func main() {
	cat := animal{"Cat", 4, false}

	fmt.Println(cat.name, "has", cat.feet, "feet.")

	if cat.hasPaws {
		fmt.Println(cat.name, "has paws.")
	}

	var dog = new(animal)

	dog.name = "Dog"
	dog.feet = 4
	dog.hasPaws = true

	fmt.Println(dog.name, "has", dog.feet, "feet.")

	if dog.hasPaws {
		fmt.Println(dog.name, "has paws.")
	}
}
