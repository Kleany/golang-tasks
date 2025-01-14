package main

import f "figures"

func main() {
	var (
		rectangle f.Rectangle = f.NewRectangle(2., 3.)
		circle    f.Circle    = f.NewCircle(4.)
	)

	f.PrintFigureInfo(&rectangle)
	f.PrintFigureInfo(&circle)
}
