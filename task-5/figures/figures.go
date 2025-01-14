package figures

import "math"

// Base struct "figure"
type figure struct {
	name string
	area float64
}

func (r *figure) Area() float64 {
	return r.area
}

func (r *figure) Type() string {
	return r.name
}

// Composed struct "Rectangle"
type Rectangle struct {
	figure
	a float64
	b float64
}

func NewRectangle(a float64, b float64) Rectangle {
	return Rectangle{
		a: a,
		b: b,
		figure: figure{
			name: "Прямоугольник",
			area: a * b,
		},
	}
}

// Composed struct "Circle"
type Circle struct {
	figure
	r float64
}

func NewCircle(r float64) Circle {
	return Circle{
		r: r,
		figure: figure{
			name: "Окружность",
			area: math.Pi * r * r,
		},
	}
}
