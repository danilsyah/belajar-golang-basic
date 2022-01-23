package main

import "fmt"

func AddMultipleArray(value ...int) {
	fmt.Println(value)

	var arrayMulti [2][2]int
	k := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			arrayMulti[i][j] = value[k]
			k++
		}
	}

	fmt.Println(arrayMulti)

}

func BilArray(n ...int) {
	var sum int
	var newSlice = make([]int, 0, 10)
	// var resultArr = make([]int, 5, 10)

	for i := 0; i < len(n); i++ {
		for j := i; j < len(n); j++ {
			sum = n[i] + n[j]
			if sum == 7 {
				newSlice = append(newSlice, n[i])
				newSlice = append(newSlice, n[j])

			}
		}
	}
	AddMultipleArray(newSlice...)

	// arrayOfArray[0][0] = 5
	// arrayOfArray[0][1] = 2
	// arrayOfArray[1][0] = -4
	// arrayOfArray[1][1] = 11
	// fmt.Println(arrayOfArray)
}

func main() {
	bil := []int{3, 5, 2, -4, 8, 11}

	BilArray(bil...)

}
