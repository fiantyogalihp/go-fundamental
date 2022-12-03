package main

// import "fmt"

// type Address struct {
// 	City, Province, Country string
// }

// /*
//  * pointer pada function sebagai parameter

//  ? peletakan pointer ada pada type data struct, menandakan type data tersebut harus pointer

//  * Tips : jika kita menggunakan data yang me-query field besar, katakanlah 50 field, pada setiap pemanggilan func akan menyebabkan memory bengkak akibat pass by value, maka usahakan menggunakan pointer
// */
// func changeCountryToIndonesia(address *Address) {
// 	address.Country = "Indonesia"
// }

// func main() {
// 	address1 := Address{
// 		City:     "Malang",
// 		Province: "Jawa Timur",
// 		Country:  "Indonesia",
// 	}
// 	// ? jika kita menggunakan pointer `&` pada variable baru, maka variable tersebut akan mengacu pada data yang sama dengan variable asal pada memory, ini yang dinamakan pass by reference
// 	var address2 *Address = &address1 // * jadi address2 adalah pointer
// 	var address3 *Address = &address1

// 	// * mengubah struct field, address1 akan ikut berubah karena address2(sebagai pointer) merupakan reference dari data address1
// 	address2.City = "Palembang"

// 	fmt.Println("mengubah struct field:")
// 	fmt.Println(address1) // ! address1 ikut berubah
// 	fmt.Println(address2)

// 	/*
// 		* mengubah seluruh struct field menggunakan pointer variable [`&` diikuti dengan data dari struct Address]

// 		? ketika kita melakukan perubahan(assign ulang), maka akan membuat data baru dari data reference nya
// 	*/
// 	// address2 = &Address{City: "Semarang", Country: "indonesia", Province: "jawa tengah"}
// 	// fmt.Println("mengubah seluruh data struct menggunakan pointer variable:")
// 	// fmt.Println(address1) // ! address1 tidak akan berubah
// 	// fmt.Println(address2)

// 	/*
// 	 * mengubah seluruh data struct dan membuat data baru dari data reference nya

// 	 * sekaligus memaksa variable lain yg mengacu pada data reference yang sama, berubah menjadi data ointer baru yang dibuat menggunakan pointer variable dengan menggunakan operator `*`
// 	 */
// 	*address2 = Address{City: "bandung", Country: "indonesia", Province: "jawa barat"}

// 	// ! data dari address1 dan address2 yang awalnya mempunyai reference sendiri akan dipaksa mengacu pada reference ke address2 ketika pointer address2 menggunakan operator `*`
// 	fmt.Println(address1)
// 	fmt.Println(address2)
// 	fmt.Println(address3)

// 	// ? cara membuat pointer kosong menggunakan func `new()`
// 	var address4 *Address = new(Address)
// 	address5 := address4
// 	address5.City = "Makassar"

// 	fmt.Println(address4) // * address4 ikut berubah, karena address5 adalah pointer
// 	fmt.Println(address5)

// 	var alamat Address = Address{City: "Panjen", Province: "jawa timur", Country: ""} // ? similar to alamat := Address{City: "Panjen", Province: "jawa timur", Country: ""}
// 	changeCountryToIndonesia(&alamat)                                                 // * pada isi parameter func pointernya, data akan diisi dengan type data yang menandakan pointer
// 	fmt.Println(alamat)
// }
