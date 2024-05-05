// NOTE THIS IS A REALLY BAD NAME FOR A MODULE
package imports

import "fmt"

type Ticket struct {
	ID    int
	Event string
}

func (t Ticket) PrintEvent() {
	fmt.Println(t.Event)
}
