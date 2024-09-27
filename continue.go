package main

import "fmt"

func ContinueModule() {
	for i := 0; i < 10; i++ {
		if 1%2 == 0 {
			continue
			// * fungsi continue dalah untuk meng-skip block code dibawahnya, dan melanjutkan ke perulangan selanjutnya
		}
		fmt.Println("perulangan ke-", i)
	}
}
