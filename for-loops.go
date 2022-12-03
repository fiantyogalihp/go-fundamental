package main

// import "fmt"

// func main() {
// 	// * opsi 1 for loops
// 	// counter := 1

// 	// for counter <= 10 {
// 	// 	fmt.Println("perulangan ke-", counter)
// 	// 	counter++
// 	// }

// 	// * opsi 2 for loops
// 	// for counter := 1; counter <= 10; counter++ {
// 	// 	fmt.Println("perulangan ke-", counter)
// 	// }
// 	// ? 'counter := 1' merupakan init statement, yang dijalankan hanya sekali saja
// 	// ? sedangkan, 'counter++' merupakan post statement, yang dijalankan setiap kali block code for selesai dijalankan
// 	// ? serta tanda ';' sebagai pembatas antara ini init statement, kondisi yang dibuat, dan post statement. perintah dijalankan berurutan dari kiri ke kanan

// 	names := []string{"fiantyo", "nduta", "zidan", "zakik", "seprin"} // * slice

// 	// slice := make([]string, 5, 7)
// 	// slice = []string{
// 	// 	"fiantyo", "nduta", "zidan", "zakik", "seprin",
// 	// }

// 	for i := 0; i < len(names); i++ {
// 		fmt.Println(names[i])
// 	}

// 	// * for range

// 	// for i, value := range slice {
// 	// 	fmt.Println("index", i, "=", value)
// 	// }
// 	// * jika variable index tidak dibutuhkan / tidak digunakan, maka ganti varibale tersebut dengan tanda underscore [ _ ], example:
// 	for _, value := range names {
// 		fmt.Println("name =", value)
// 	}

// 	person := make(map[string]string)

// 	person["name"] = "fiantyo"
// 	person["hobby"] = "animehh"
// 	person["title"] = "programmer"

// 	for key, value := range person {
// 		fmt.Println(key, "=", value)
// 	}
// }
