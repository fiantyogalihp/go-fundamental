package main

import "fmt"

func random() interface{} {
	return "wuuuyyyy"
	// * return value string
}

func TypeAssertionModule() {
	/*
		! type assertions hanya bisa dilakukan ketika yakin dengan type data dari kembalian nya sama dengan type data yang akan diubah, jika tidak maka akan terjadi panic pada program
	*/
	var result interface{} = random()
	// * return convert to string
	// var resultString = result.(string)
	// fmt.Println(resultString)

	// * agar lebih aman ketika menggunakan type assertions, kita bisa menggunakan type assertions switch
	switch value := result.(type) {
	case string:
		fmt.Println("string:", value, "is string")
	case int:
		fmt.Println("int:", value, "is int")
	case bool:
		fmt.Println("bool:", value, "is bool")
	default:
		fmt.Println("Unknown type")
	}
}
