package main

import "fmt"

func ConstantModule() {
	// const name string = "fian"
	// const age = 18
	// const country = "indonesia"

	const (
		name    string = "fian"
		age     int    = 18
		country string = "indonesia"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

}
