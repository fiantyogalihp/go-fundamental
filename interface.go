package main

// import "fmt"

// /*
// 	? dalam interface, meskipun structnya berbeda, tetap bisa memakai kontrak(interface) yang sama
// */
// type HasName interface {
// 	getName() string
// }

// // * struct 1
// type Person struct {
// 	Name string
// }

// func (person Person) getName() string {
// 	return person.Name
// }

// // * struct 2
// type Animal struct {
// 	Name string
// }

// func (animal Animal) getName() string {
// 	return animal.Name
// }

// // * function general
// func sayHello(hasName HasName) {
// 	fmt.Println("hello", hasName.getName())
// }

// /*
// 	? Interface Kosong(empty)

// 	? interface kosong ini bisa menyimpan type data apapun seperti hal nya fmt.Println() yang bisa mengembalikan type data apapun

// 	? cara membuat interface kosong ini cukup menggunakan type data `interface{}`
// */
// func sayHi(i int) interface{} {
// 	if i == 1 {
// 		return 1
// 	} else if i == 2 {
// 		return false
// 	} else {
// 		return "anjayy"
// 	}
// }

// func main() {
// 	budi := Person{Name: "budi"}

// 	sayHello(budi)

// 	cat := Animal{Name: "Kucing"}

// 	sayHello(cat)

// 	// * interface kosong
// 	var data interface{} = sayHi(1) // * data := sayHi(1)
// 	fmt.Println(data)
// }
