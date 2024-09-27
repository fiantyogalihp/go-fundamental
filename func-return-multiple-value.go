package main

import (
	"fmt"
)

func getFullName(firstName string, lastName string) (string, string) {
	if firstName == "" || lastName == "" {
		return "Hello", "cuy"
	}

	return firstName, lastName
}

// func getFullName() (string, string) {
// 	return "budi", "kopleng"
// }

func FuncReturnMultipleValueModule() {
	firstName, lastName := getFullName("Fiantyo", "kopleng")
	// * cara menghiraukan return value function jika tidak diperlukan dengan tanda [ _ ]
	// firstName, _ := getFullName()

	fmt.Println(firstName)
	fmt.Println(lastName)
}
