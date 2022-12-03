package helper

import "fmt"

var gilang = "gilang pedo" // ! tidak bisa di akses dari luar package
/*
 ! access modifier pada golang sangat simple, ketika kita menggunakan awal huruf kecil, maka tidak akan bisa diakses package lain,
 ! sebaliknya ketika menggunakan awal huruf kapital, maka func/variable tersebut bisa diakses package lain
*/
func SayHello(name string) {
	fmt.Println("Hello", name, ", tentunya", gilang)

}

var IsPedo = true // * bisa di akses dari luar package
