package main

import (
	"fmt"
	"seachUser/lib"
	"strings"
)



func main() {
	dataName := []string{"ari", "eka", "saputra", "jhon"}
	var input string

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&input)

	input = strings.ToLower(input)

	lib.SearchPerson(dataName, input)
}