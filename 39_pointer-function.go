package main

import "fmt"

type Alamat struct {
	Kota, Provinsi, Negara string
}

func ChangeCountry(address *Alamat) {
	address.Negara = "Indonesia"
}

func main() {
	var address = Alamat{"Subang", "Jawa Barat", "Inggris"}
	var alamat = &address
	ChangeCountry(alamat)
	fmt.Println(alamat)
}
