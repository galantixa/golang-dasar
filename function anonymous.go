package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "Admin"
	}
	registerUser("Admin", blackList)
	registerUser("Fajar", blackList)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("Fajar", func(name string) bool {
		return name == "root"
	})
}
