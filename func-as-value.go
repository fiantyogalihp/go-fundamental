package main

import "fmt"

func getGoodBye(name string) string {
	return "hello " + name
}

func FuncAsValueModule() {
	// ? memasukan func ke dalam variable dan menjadikannya sebagai value
	sayGoodBye := getGoodBye

	result := sayGoodBye("bapak budi")
	fmt.Println(result)

}
