package main

import "fmt"

func IfStatementModule() {
	name := "fian"

	if name == "fian" {
		fmt.Println("hello", name)
	} else if name == "slamet" {
		fmt.Println("hello", name)
	} else {
		fmt.Println("iyahhh")
	}

	// * Short statement pada if
	// * opsi biasa
	// length := len(name)
	// if length > 5 {
	// 	fmt.Println("Nama terlalu panjang")
	// } else {
	// 	fmt.Println("Nama sudah benar")
	// }

	// * opsi short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
