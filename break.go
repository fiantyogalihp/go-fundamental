package main

import "fmt"

func BreakModule() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
			// * fungsi break adalah untuk menghetikan semua perulangan
		}
		fmt.Println("perulangan ke-", i)
	}
}
