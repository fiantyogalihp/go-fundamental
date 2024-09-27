package main

import "fmt"

type Blacklist func(name string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are BLOCKED!", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func AnonymousFuncModule() {
	blacklist := func(name string) bool {
		if name == "anjing" || name == "babi" {
			return name == "..."
		}
		return name == "root" || name == "admin"
	}

	registerUser("budi", blacklist)
	registerUser("anjing", blacklist)

	registerUser("Eko", func(name string) bool {
		return name == "Eko"
	})
}
