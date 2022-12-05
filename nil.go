package main

// import "fmt"

/*
 ? nil hanya bisa digunakan di beberapa type data seperti interface, func, map, slice, pointer, dan channel
*/

// type newNameMap map[string]string

// func newMap(name string) newNameMap {
// 	if name == "" {
// 		return nil
// 	} else {
// 		return newNameMap{
// 			"Name": name,
// 		}
// 	}
// }

// func main() {
// 	slamet := newMap("Slamet")
// 	fmt.Println(slamet)

// 	var budi newNameMap = newMap("Budi")
// 	// ? jika tidak menggunakan nil sebagai compare expression, maka tidak efektif seperti ini if budi["name"] == ""
// 	if budi == nil {
// 		fmt.Println("Data kosong")
// 	} else {
// 		fmt.Println(budi)
// 	}
// }
