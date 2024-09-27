package main

import (
	"flag"
	"fmt"
)

func FlagModule() {
	fmt.Println("How to use: go run flag.go -host=YOUR_DATABASE_HOST -user=YOUR_DATABASE_USER -password=YOUR_DATABASE_PASSWORD")

	host := flag.String("host", "localhost", "-host=YOUR_DATABASE_HOST")
	username := flag.String("user", "root", "-user=YOUR_DATABASE_USER")
	password := flag.String("password", "", "-password=YOUR_DATABASE_PASSWORD")

	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*username)
	fmt.Println(*password)
}
