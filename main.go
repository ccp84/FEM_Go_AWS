package main

import (
	"femBasics/circle"
	"femBasics/imports"
	"fmt"
)

func main() {
	newTicket := imports.Ticket{
		ID:    1,
		Event: "FEM Workshop",
	}
	newTicket.PrintEvent()

	myCircle := circle.NewCircle(5)

	fmt.Println("The circumference of my circle is: ", myCircle.Circumference())
	fmt.Println("The area of my circle is: ", myCircle.Area())
}
