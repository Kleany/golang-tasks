package main

import (
	"fmt"
)

func main() {
	var strings = []string{"aaa", "bbb", "ccc", "ddd", "eee"}

	for i := 0; i < len(strings)-1; i++ {
		fmt.Printf("Comparing strings \"%v\" and \"%v\":\t", strings[i], strings[i+1])

		if strings[i] > strings[i+1] {
			fmt.Printf("\"%v\" < \"%v\"\n", strings[i], strings[i+1])
			fmt.Println("False")
			return
		} else {
			fmt.Printf("\"%v\" > \"%v\"\n", strings[i], strings[i+1])
		}
	}

	fmt.Println("True")
}
