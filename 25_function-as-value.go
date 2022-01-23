package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {

	// function sebagai nilai variabel
	sayGoodBye := getGoodBye

	result := sayGoodBye("danil")

	fmt.Println(result)
}
