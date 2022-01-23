package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noktp NoKTP = "32121291291291"
	var status Married = true

	fmt.Println(noktp)
	fmt.Println(status)
}
