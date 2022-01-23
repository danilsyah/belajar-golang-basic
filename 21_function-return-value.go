package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello bro"
	} else {
		return "Hello " + name
	}
}

func getHasil(nilai int) string {

	var keterangan string

	if nilai >= 70 {
		keterangan = "Lulus"
	} else {
		keterangan = "Tidak Lulus"
	}

	return keterangan
}

func main() {
	result := getHello("danil")

	fmt.Println(result)

	fmt.Println(getHello(""))

	fmt.Println(getHasil(70))
}
