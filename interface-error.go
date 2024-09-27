package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func InterfaceErrorModule() {
	result, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Result =", result)
	} else {
		// ? hasil kembalian error
		fmt.Println("Error:", err.Error())
	}
}
