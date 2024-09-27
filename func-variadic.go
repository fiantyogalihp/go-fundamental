package main

import "fmt"

func sumAll(numbers ...int) (total int) {
	total = 0
	for _, v := range numbers {
		total += v
	}
	return
}

func FuncVariadicModule() {
	result := sumAll(10, 15, 20, 20)
	fmt.Println("total :", result)

	// ? memasukan data yang sudah tersedia ke dalam variadic parameter
	slice := []int{10, 10, 90, 35}
	result = sumAll(slice...)
	fmt.Println(result)
}
