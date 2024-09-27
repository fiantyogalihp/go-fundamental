package main

import "fmt"

/*
	? fungsi recover, yaitu menangkap data panic dan tetap menjalankan program berjalan yang dihentikan oleh panic(dalam func)

	! penempatan posisi recover harus diluar func dimana panic dijalankan, karena panic akan menghentikan semua proses program yang berjalan dalam func tersebut, solusinya adalah menempatkan recover dalam defer func
	* recover yang salah, yaitu ketika menjadikan satu tempat dengan panic
*/
func endAppV2() {
	message := recover()
	if message != nil {
		fmt.Println("ERROR with message: ", message)
	}

	fmt.Println("aplikasi selesai")
}

func runAppV2(error bool) {
	defer endAppV2()
	if error {
		panic("APLIKASI ERROR") // ! panic akan menghentikan program berjalan(dalam func)
	}

	fmt.Println("aplikasi berjalan")
}

func RecoverModule() {
	runAppV2(true)
	fmt.Println("Fiantyo") // ? block code ini akan tetap dijalankan ketika panic sudah di recover dalam runApp func
}
