package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

/*
	? Struct func ini sebenarnya merupakan func standalone, Struct tidak mempunyai behavior sepertti class yang mempunyai func sendiri,

	? struct func dibuat agar spesifik untuk kebutuhan Struct itu sendiri. Jadi seakan akan Struct func ini mempunyai func(behavior)
*/
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello,", name, "my Name is", customer.Name)
}

func StructModule() {
	// * Cara ke-1
	// var eko Customer
	// eko.Name = "Fiantyo"
	// eko.Address = "Stasiun Ngebruk street"
	// eko.Age = 18

	// fmt.Println(eko)
	// fmt.Println(eko.Name)
	// fmt.Println(eko.Address)
	// fmt.Println(eko.Age)

	// * Cara ke-2
	budi := Customer{
		Name:    "Budi",
		Address: "Kampung warna-warni",
		Age:     20,
	}

	fmt.Println(budi)

	// * Cara ke-3
	/*
		! cara ke-3 ini tidak disarankan, karena ketika field/struktur data dari Struct itu sendiri dirubah, maka akan rentan(rawan) error, karena data menyesuaikan urutan dari Struct itu sendiri
	*/
	// slamet := Customer{"slamet", "kampung kopleng", 20}
	// fmt.Println(slamet)

	/*
	 ? Struct Method

	 * cara pemanggilan Struct ini dari Struct yang telah dideklarasikan dan mempunyai value, lalu diberi func(behavior) sesuai Struct func yang telah dibuat, seperti dibawah ini
	*/
	budi.sayHello("Slamet")
}
