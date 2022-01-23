package main

import "fmt"

func getFullName(firstName, middleName, lastName string) (string, string, string) {
	return firstName, middleName, lastName
}

func main() {
	namaDepan, _, namaAkhir := getFullName("danil", "syah", "arihardjo")

	fmt.Println(namaDepan)
	// fmt.Println(namaTengah)
	fmt.Println(namaAkhir)
}
