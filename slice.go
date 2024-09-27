package main

import "fmt"

func SliceModule() {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	slice1 := months[5:8]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // * menhitung jumlah array pada slice
	fmt.Println(cap(slice1)) // * menghitung jumlah kapasitas dari slice

	// ! slice adalah sepenuhnya bergantung pada array, jika value array dirubah maka value slice juga berubah
	// // * merubah array
	// months[5] = "Juni"
	// fmt.Println(slice1)

	// // * merubah slice
	// slice1[2] = "Agustus"
	// fmt.Println(months)

	slice2 := months[10:]
	fmt.Println(slice2)
	slice2[1] = "bukan Desember"
	fmt.Println(slice2)

	slice3 := append(slice2, "isekai") // * membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas penuh maka akan membuat array baru dan array lama tidak akan berubah, jika sebaliknya maka data baru akan me-replace data index yang sudah ada
	fmt.Println(slice3, "(array baru)")
	fmt.Println(months)

	// ! hati-hati saat merubah nilai array, karena akan berdampak pada semua slice yang mengacu pada array tersebut

	newSlice := make([]string, 2, 7) // * opsi paling aman ketika membuat slice karena array yang dipakai disembunyikan dibelakang slice

	newSlice[0] = "fian"
	newSlice[1] = "tyo"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	/*
	 * Bisa juga dideklarasikan menggunakan make dengan kapasitas 0
	 * Bedanya adalah bagaimana golang membandingkan itu dengan nil
	 */
	var nilSlice []int // * empty slice
	emptySlice1 := make([]int, 0)
	emptySlice2 := []int{}

	fmt.Println(nilSlice == nil)    // true
	fmt.Println(emptySlice1 == nil) // false
	fmt.Println(emptySlice2 == nil) // false

	// * Copy slice from another slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	// ? PERBEDAAN ARRAY DENGAN SLICE
	iniArray := [5]int{1, 2, 3, 4, 5} // * ini Array
	iniSlice := []int{1, 2, 3, 4, 5}  // * ini Slice

	// * perbedaan nya jika array harus ada lenght (nilai tetap / ... ), sedangkan slice dibiarkan kosong

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
