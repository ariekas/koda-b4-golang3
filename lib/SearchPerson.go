package lib

import (
	"fmt"
)

func SearchPerson(dataName []string, input string) {
	for _, value := range dataName {
		if value == input {
			fmt.Println("User ditemukan:", input)
			break
		} else {
			panic(fmt.Sprintf("user '%s' tidak di temukan!", input))
		} 
	}
}