package main

// import (
// 	"container/list"
// 	"fmt"
// )

// func main() {
// 	/*
// 	 ! Data list pada awal dan akhir adalah nil
// 	 ! Data list tidak mempunyai index, jadi iterasi 1 by 1 dengan kondisi berhenti nil
// 	*/
// 	data := list.New()

// 	data.PushBack("Fiantyo")
// 	data.PushBack("Galih")
// 	data.PushBack("Pramudya")

// 	data.PushFront("Slamet")

// 	fmt.Println("mengambil data paling depan:", data.Front())
// 	fmt.Println("mengambil data paling belakang:", data.Back())

// 	fmt.Println("mengambil data setelah nya dari depan:", data.Front().Next())
// 	fmt.Println("mengambil data sebelumnya dari belakang:", data.Back().Prev())

// 	fmt.Println("mengurutkan data list dari depan ")
// 	for element := data.Front(); element != nil; element.Next() {
// 		fmt.Println(element.Value)
// 	}

// 	fmt.Println("mengurutkan data list dari belakang")
// 	for element := data.Back(); element != nil; element.Prev() {
// 		fmt.Println(element.Value)
// 	}
// }
