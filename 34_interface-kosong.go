package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

// interface sebagai tipe data parameter dan return
func Test(value interface{}) interface{} {
	return value
}

func main() {
	var angka interface{} = Ups(3)
	fmt.Println(angka)

	var nama interface{} = Test("danil")
	fmt.Println(nama)

	var tahunLahir interface{} = Test(1994)
	fmt.Println(tahunLahir)

	var menikah interface{} = Test(false)
	if menikah == true {
		fmt.Println("menikah")
	} else {
		fmt.Println("Single")
	}
}
