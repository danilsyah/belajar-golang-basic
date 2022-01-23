package main

import "fmt"

// membuat struct
type Customer struct {
	Name, Address string
	Age           int
}

// struct method / function
func (customer Customer) sayHi(name string) {
	fmt.Println("hai", name, "My name is ", customer.Name)
}

func main() {
	// deklarasi nilai struct
	var danil Customer
	danil.Name = "Danil"
	danil.Address = "Bintan"
	danil.Age = 27

	danil.sayHi("udin")

	// fmt.Println(danil.Name)

	// haykal := Customer{
	// 	Name:    "Haykal",
	// 	Address: "Sumedang",
	// 	Age:     3,
	// }

	// fmt.Println(haykal.Name)

	// if danil.Age == haykal.Age {
	// 	fmt.Println("Bapak")
	// } else {
	// 	fmt.Println("Anak")
	// }

}
