package main

import "fmt"

func random() interface{} {
	return "danil"
}

func main() {
	var result interface{} = random()
	// mengkonversi dari tipe data interface ke string
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is Int")
	case bool:
		fmt.Println("Value", value, "is Bool")
	default:
		fmt.Println("Unknown type")
	}
}
