package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("First test\n")
	fmt.Print(diceRole())
}

func diceRole() int {

	return rand.Intn(6)
}
