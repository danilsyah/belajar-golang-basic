package main

import "fmt"

func main() {
	name := "danill"

	if name == "danil" {
		fmt.Println("nama cocok")
	} else {
		fmt.Println("nama tidak cocok")
	}

	// if short statement
	if result := 5 + 10; result > 10 {
		fmt.Println("jumlah lebih dari 10")
	} else {
		fmt.Println("jumlah kurang dari 10")
	}
}
