package main

import (
	"fmt"
	"os"
	"seachUser/lib"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi kesalahan:", r)
			fmt.Println("Program dihentikan.")
			os.Exit(1)
		} 
	}()

	dataName := []string{"ari", "eka", "saputra", "jhon"}
	var input string
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&input) 

	newInput := &input
	*newInput = strings.ToLower(*newInput)

	lib.SearchPerson(dataName, *newInput) 
}