package main

import "fmt"

func ComparisonOperator() {
	var name1 = "fian"
	var name2 = "fian"

	var result = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}
