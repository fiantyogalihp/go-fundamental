package main

import "fmt"

func endApp() {
	fmt.Println("aplikasi selesai")
}

func runApp(error bool) {
	defer endApp() // ? defer akan tetap dijalankan meskipun panic menghentikan program berjalan(dalam func)
	if error {
		panic("APLIKASI ERROR") // ? pesan panic akan ditampilkan setelah defer ditampilkan
		// ! panic akan menghentikan program berjalan(dalam func)
	}
	fmt.Println("aplikasi berjalan")
}

func PanicModule() {
	runApp(true)
}
