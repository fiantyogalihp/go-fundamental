package main

import (
	"fmt"
	"strconv"
)

func StrconvModule() {
	fmt.Println("convertion dari string ke type data lain:")
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	number, err := strconv.ParseInt("120000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("Error", err.Error())
	}

	float, err := strconv.ParseFloat("1.90", 64)
	if err == nil {
		fmt.Println(float)
	} else {
		fmt.Println("Error", err.Error())
	}

	parseInt, _ := strconv.Atoi("2000000")
	fmt.Println(parseInt)

	fmt.Println("convertion dari type data lain ke string:")
	formatInt := strconv.FormatInt(1550000, 10)
	fmt.Println(formatInt)

	formatINt := strconv.Itoa(1234567)
	fmt.Println(formatINt)
}
