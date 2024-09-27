package main

import "fmt"

// ? return named value bisa menamai return value menggunakan variable, data bisa complex, dan bisa menjadi pengembalian data tanpa mendeklarasika ulang
func getFullNameV2() (firstName, middleName, lastName string) {
	firstName = "fiantyo"
	middleName = "galih"
	lastName = "pramudya"

	// return firstName, middleName, lastName // * mendeklarasikan ulang return named value tidak wajib dan bisa dihilangkan
	return
}

func getAddress() (villageName string, RT, RW int) {
	villageName = "ngebruk"
	RT = 19
	RW = 03

	return
}

func FuncReturnNamedValue() {
	a, b, c := getFullNameV2()
	villageName, _, _ := getAddress()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(villageName)
}
