package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "danil",
		"address": "sumedang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])

	person["title"] = "programmer"

	fmt.Println(person)

	delete(person, "title")

	fmt.Println(person)

	// membuat map baru
	book := make(map[string]string)

	book["title"] = "Buku Golang"
	book["author"] = "Danil"
	book["wrong"] = "ups"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}
