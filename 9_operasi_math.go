package main

import "fmt"

func main() {

	// operator aritmatika
	var a = 10
	var b = 3
	var tambah = a + b
	var kurang = a - b
	var kali = a * b
	var bagi = a / b
	var modulus = a % b

	fmt.Println(tambah)
	fmt.Println(kurang)
	fmt.Println(kali)
	fmt.Println(bagi)
	fmt.Println(modulus)

	// augmented assignments
	// += , -=, *=, /= , %=
	var i = 10
	i += 10
	fmt.Println(i)

	// unary operator
	// i++, i--, ! , -
	var nilaiNegatif = -100
	var nilaiPositif = +100
	var nikah = false
	var j = 10
	j++
	fmt.Println(nilaiNegatif)
	fmt.Println(nilaiPositif)
	fmt.Println(j)
	fmt.Println(!nikah)
}
