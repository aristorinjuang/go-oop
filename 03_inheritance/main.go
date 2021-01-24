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

func main() {
	dragon := monster{
		animal{"Dragon", 4, true},
		[]string{"Nuclear blast", "Ice smoke"},
	}

	dragon.print()
}
