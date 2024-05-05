package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// capitalised function can be made public
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

// adding the Struct to the front makes this a method of that type
func (p *Person) changeName(newName string) {
	p.Name = newName
}

func main() {

	myPerson := Person{
		Name: "Me",
		Age:  2,
	}
	// +v adds the interface to the output
	fmt.Printf("%+v\n", myPerson)

	myPerson.Name = "Changed"

	fmt.Printf("%+v\n", myPerson)

	// -------------------------- Using additional functions from here

	newPerson := NewPerson("New", 5)

	newPerson.changeName("Changed")

	fmt.Printf("This is my new person %+v\n", newPerson)

	// Pointers and references

	a := 7
	b := &a
	// b = 9  will not work as it is a pointer to a
	// dereference and add a new pointer instead
	// this also changes a as b is pointing to the memory slot of a
	*b = 9

	fmt.Println(*b)
	fmt.Println(a)

	mySlice := []int{
		1, 2, 3,
	}
	for index := range mySlice {
		mySlice[index]++
	}
	fmt.Println(mySlice)

}
