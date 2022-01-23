package main

import "fmt"

func main() {
	// >, <, >=, <=, ==, !=

	var name1 = "danil"
	var name2 = "danil"
	var result bool = name1 == name2
	fmt.Println(result)

	var nilai1 = 200
	var nilai2 = 100

	fmt.Println(nilai1 > nilai2)
	fmt.Println(nilai1 < nilai2)
	fmt.Println(nilai1 >= nilai2)
	fmt.Println(nilai1 <= nilai2)
	fmt.Println(nilai1 == nilai2)
	fmt.Println(nilai1 != nilai2)

}
