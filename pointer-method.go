package main

import "fmt"

type Man struct {
	Name string
}

/*
 * diusahakan ketika kita menggunakan struct method, kita perlu memakai pointer method didalamnya, untuk 	menghindari pembengkaan memory
 */
//
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func PointerMethodModule() {
	budi := Man{"Budi"}
	budi.Married()

	fmt.Println(budi)
}
