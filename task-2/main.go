package main

import (
	"fmt"
)

func main() {
	var nums = [7]int8{1, 2, 3, 4, 5, 6, 1}

	var valuesCount = map[int8]int8{}

	for _, value := range nums {
		valuesCount[value]++

		if valuesCount[value] == 2 {
			fmt.Println("True")
			return
		}
	}

	fmt.Println("False")
}
