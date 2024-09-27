package main

import (
	"fmt"
)

// * membuat Alias function ketika menjadikan func-as-parameter agar  mendeklarasikan sebagai parameter tidak terlalu panjang dan bisa di baca dengan mudah.
// * serta jika ada perubahan pada suatu function maka kita cuku merubah nya pada type declaratin
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

// ! tidak disarankan menggunakan filter as parameter di dalam func utama filter, karena akan memperbanyak opsi dan variable filter yang akan dipakai, jadi kita memfokuskan fitur filter dalam func uatama
func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func FuncAsParameterModule() {
	sayHelloWithFilter("Budi", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Slamet", filter)
}
