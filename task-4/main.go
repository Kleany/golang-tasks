package main

import (
	"fmt"
)

func main() {
	var myString = "съешь ещё этих мягких французских булок, да выпей чаю"

	var symbolsCount = map[string]int8{}

	for _, symbol := range myString {
		symbolsCount[string(symbol)]++
	}

	for key, value := range symbolsCount {
		fmt.Printf("Symbol: \"%v\"\tCount: %v\n", key, value)
	}
}
