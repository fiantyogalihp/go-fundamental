package main

import "fmt"

func SwitchModule() {
	name := "fian"

	switch name {
	case "fian":
		fmt.Println("hello", name)
	case "budi":
		fmt.Println("hello", name)
	default:
		fmt.Println("iyahhhh")
	}

	// * short statement
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// * switch tanpa kondsi (if statement dalam bentuk sederhana)
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama terlalu panjang")
	}
}
