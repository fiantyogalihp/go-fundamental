package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func ContainerRingModule() {
	// ! data container ring dimulai dari 0
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "mas wahyu ke-" + strconv.Itoa(i+1)
		data = data.Next()
	}

	// fmt.Println(*data) // ! data container ring tidak bisa di println, makanya untuk println datanya menggunakan func bawaan dari container ring
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
