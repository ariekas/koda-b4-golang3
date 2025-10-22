package lib

import (
	"fmt"
)

func SearchPerson(dataName []string, input string) {
	for i := range len(dataName){
		if dataName[i] == input {
			fmt.Println("User ditemukan:", input)
			break
		}else{
			fmt.Println([]string{})
			break
		}
	}
}