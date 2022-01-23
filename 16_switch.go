package main

import "fmt"

func main() {
	name := "danil"

	switch name {
	case "danil":
		fmt.Println("Hello danil")
	case "haykal":
		fmt.Println("Hello haykal")
	default:
		fmt.Println("Kenalan dong")
	}

	nilai := 50
	var grade string

	switch {
	case nilai >= 90:
		grade = "A"
	case nilai >= 80:
		grade = "B"
	case nilai >= 70:
		grade = "C"
	default:
		grade = "E"
	}

	fmt.Println(grade)

	// short switch
	switch hasil := grade; hasil == "E" {
	case true:
		fmt.Println("anda tidak lulus")
	case false:
		fmt.Println("anda lulus")
	}
}
