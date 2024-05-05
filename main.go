package main

import (
	"femBasics/imports"
)

func main() {
	newTicket := imports.Ticket{
		ID:    1,
		Event: "FEM Workshop",
	}
	newTicket.PrintEvent()
}
