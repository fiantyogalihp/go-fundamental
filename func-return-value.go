package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Cuy!"
	}

	return "Hello " + name
}

func FuncReturnValueModule() {
	result := getHello("Fiantyo Galih")

	fmt.Println(result)
	fmt.Println(getHello(""))
}
