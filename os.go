package main

import (
	"fmt"
	"os"
)

func OSModule() {
	args := os.Args
	fmt.Println(args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
