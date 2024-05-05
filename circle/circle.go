package circle

const (
	PI = 3.14
)

type Circle struct {
	Radius float64
}

func NewCircle(r float64) Circle {
	return Circle{Radius: r}
}

func (c Circle) Circumference() float64 {
	return 2 * (c.Radius * PI)
}

func (c Circle) Area() float64 {

	return PI * (c.Radius * c.Radius)
}
