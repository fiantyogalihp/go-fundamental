package main

import "fmt"

func VariablesModule() {
	// * Deklarasi awal variable, jika untuk perubahan selanjutnya cukup langsung menggunakan nama variablenya saja

	// * contoh 1
	var name string

	name = "bapak budi"
	fmt.Println(name)

	name = "bapak samad"
	fmt.Println(name)

	// * contoh 2
	var age = 24
	fmt.Println(age)

	age = 25
	fmt.Println(age)

	// * contoh 3
	country := "Indonesia"
	fmt.Println(country)

	country = "Singapore"
	fmt.Println(country)

	// * Multiple variable
	var (
		firstName = "Fian"
		lastname  = "tyo"
	)

	fmt.Println(firstName)
	fmt.Println(lastname)

}
