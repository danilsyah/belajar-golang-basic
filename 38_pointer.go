package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Sumedang", "Jawa Barat", "Indonesia"}
	var address3 = &address1
	fmt.Println(address1)
	fmt.Println(address3)

	var address2 *Address = &address1
	address2.City = "Bandung"
	address1.Country = "Eropa"
	fmt.Println(address2)
	fmt.Println(address1)
	fmt.Println(address3)

	*address2 = Address{"Tanjunguban", "Bintan", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address3)

	var address4 = new(Address)
	address4.City = "Wado"
	address4.Province = "jabar"
	address4.Country = "perancis"

	fmt.Println(address1)
	fmt.Println(address4)

}
