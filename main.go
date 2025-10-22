package main

import (
	"fmt"
	"strings"
)

func searchPerson(dataName []string, input string) {
	input = strings.ToLower(input)
	found := false

	for i := range len(dataName){
		if dataName[i] == input {
			fmt.Println("User ditemukan:", input)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("User tidak ditemukan")
	}
}

func main() {
	dataName := []string{"ari", "eka", "saputra", "jhon"}
	searchPerson(dataName, "waadwada")
}