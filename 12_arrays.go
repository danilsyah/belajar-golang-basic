package main

import "fmt"

func main() {
	// saat membuat array, harus menentukan jumlah data yang bisa ditampung oleh array tersebut
	// jumlah/daya tampung array tidak bisa bertambah setelah array dibuat

	var names [3]string

	names[0] = "danil"
	names[1] = "fika"
	names[2] = "haykal"

	fmt.Println(names)

	var values = [3]int{
		90,
		93,
		65,
	}

	fmt.Println(values)
	fmt.Println(values[1])
	fmt.Println(len(values))
	fmt.Println(len(names))

}
