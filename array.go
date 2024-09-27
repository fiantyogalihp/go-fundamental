package main

import "fmt"

// import "fmt"

func ArrayModule() {
	// * single decalrations array
	var names [3]string

	names[0] = "fiantyo"
	names[1] = "galih"
	names[2] = "pramudya"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// * multiple declarations array
	var values = [3]int{
		30,
		20,
		10,
	}

	fmt.Println(values)
	fmt.Println(len(values)) // * untuk mendapatkan panjang array pada array, meskipun array dalam kondisi null

	var exercises = [...]int{
		80,
		50,
		90,
		70,
		85,
	}

	fmt.Println(exercises)
	fmt.Println(len(exercises))

	exercises[0] = 100 // * untuk mengganti nilai pada index yg ditentukan dalam array
	fmt.Println(exercises)

}
