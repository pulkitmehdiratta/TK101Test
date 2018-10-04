package main

import (
	"fmt"
)

func main() {
	t := 10
	switch t {
	case 12:
		fmt.Println("Good morning!")
	case 10:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
