package main

import (
	"fmt"
	"slices"
)

func main() {
	// arrays
	animals := [2]string{}

	animals[0] = "cat"
	animals[1] = "dog"

	fmt.Println(animals)

	pets := [2]string{
		"marley",
		"charlie",
	}

	fmt.Println(pets)

	// slice is a flexible array
	animals2 := []string{
		"dog",
		"cat",
	}

	animals2 = append(animals2, "rabbit")
	fmt.Println(animals2)

	// use the slices package to "pop"
	// delete everything after 0-1
	animals2 = slices.Delete(animals2, 0, 1)
	fmt.Println(animals2)

	// loop over an array

	for i := 0; i < len(animals); i++ {
		fmt.Println("this is my animal", animals[i])
	}

	animals2 = append(animals2, "dog")

	// built in range of slices

	for index, value := range animals2 {
		fmt.Println(index, value)
	}
}
