package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	x := 0
	for x < 10 {
		twoDiceRoleValue()
		x++
	}
}

func twoDiceRoleValue() {
	role1 := roleDice()
	role2 := roleDice()

	if role1 == 0 || role2 == 9 {
		role1 = roleDice()
		role2 = roleDice()
	}
	fmt.Printf("%v + %v = %v\n", role1, role2, role1+role2)
}

func roleDice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7)
}
