package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

var rolesOnALine = 10
var howManyRoles = 1000

var role1Values [6]int
var role2Values [6]int
var totalRoleValues [12]int

func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("Cant seed the value")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func main() {
	count := 0
	for count < howManyRoles {
		count++
		twoDiceRoleValue(count)

		if count%rolesOnALine == 0 {
			fmt.Println()
		}
	}

	printFromRecordsArray()
}

func twoDiceRoleValue(count int) {
	role1 := roleDice()
	role2 := roleDice()
	total := role1 + role2

	recordRoleValues(role1, role2, total)
	printRoleValues(total)

}

func roleDice() int {
	return rand.Intn(6) + 1
}

func recordRoleValues(r1 int, r2 int, tr int) {
	//fmt.Printf("cnt=%v, total=%v\t", count, tr)

	role1Values[r1-1]++
	role2Values[r2-1]++
	totalRoleValues[tr-1]++
}

func printRoleValues(tr int) {
	fmt.Printf("%v\t", tr)
}

func printFromRecordsArray() {
	fmt.Println("Count of Each Dice 1 Values:")
	count := 0
	for count < 6 {
		fmt.Println(role1Values[count])
		count++
	}
	fmt.Println("Count of Each Dice 2 Values:")
	count = 0
	for count < 6 {
		fmt.Println(role2Values[count])
		count++
	}
	fmt.Println("Count of totals of each role:")
	count = 0
	for count < 12 {
		fmt.Println(totalRoleValues[count])
		count++
	}
}
