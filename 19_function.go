package main

import "fmt"

func sayHello() {
	fmt.Println("hai danil")
}

func main() {
	for i := 0; i < 5; i++ {
		sayHello()
	}
}
