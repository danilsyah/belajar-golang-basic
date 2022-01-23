package main

import "fmt"

// function variadic adalah sebuah function yang memiliki parameter yang mana parameter dapat menampung nilai lebih dari 1
// parameter variadic seperti array atau slice
// parameter variadic harus di sebelah kanan jika parameter function lebih dari 1
func sumAll(nama string, numbers ...int) (string, int) {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return nama, total
}

func showPerson(peoples ...string) {
	for _, value := range peoples {
		fmt.Println(value)
	}
}

func main() {
	nama, total := sumAll("danil", 10, 29, 31, 12, 43)
	fmt.Println(nama, total)

	dataSlice := []int{20, 21, 46, 21, 63}
	nama, total = sumAll("udin", dataSlice...)

	fmt.Println(total)

	peoples := []string{"danil", "haykal", "fika", "razil", "khalinda"}
	showPerson(peoples...)
}
