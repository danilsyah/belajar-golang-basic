package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "diubah"
	fmt.Println(slice1)
	fmt.Println(months)

	slice1[0] = "meiDiubah"
	fmt.Println(slice1)
	fmt.Println(months)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "SabtuBaru"
	daySlice1[1] = "MingguBaru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "LiburBaru")
	daySlice2[0] = "ups"
	// daySlice2[1] = "test"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "danil"
	newSlice[1] = "haykal"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	newSlice = append(newSlice, "fika")
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
