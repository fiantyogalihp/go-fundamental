package main

import "fmt"

func TypeDeclarationsModule() {
	type NoKTP string
	type married bool

	var noKtp NoKTP = "28934928348264"
	var marriedStatus married = false

	fmt.Println(noKtp)
	fmt.Println(marriedStatus)

}
