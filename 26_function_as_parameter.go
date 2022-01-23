package main

import "fmt"

type Filter func(string) string
type Authentication func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func login(username string, auth Authentication) string {
	result := auth(username)

	return result
}

func accessAuth(username string) string {
	if username == "admin" {
		return "anda login sebagai administrator"
	} else {
		return "anda login sebagai user biasa"
	}
}

func main() {
	sayHelloWithFilter("Danil", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)

	fmt.Println(login("danil", accessAuth))
	fmt.Println(login("admin", accessAuth))
}
