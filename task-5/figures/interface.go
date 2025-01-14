package figures

import "fmt"

// Interface
type Figure interface {
	Area() float64
	Type() string
}

func PrintFigureInfo(f Figure) {
	fmt.Printf("Type: %v\tArea: %v\n", f.Type(), f.Area())
}
