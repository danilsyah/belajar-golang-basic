package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slicePerson := []string{"danil", "haykal", "fika", "razil", "khalinda"}

	for i := 0; i < len(slicePerson); i++ {
		fmt.Println("Nama ke ", i, slicePerson[i])
	}

	for i, value := range slicePerson {
		fmt.Println("Index ke ", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Danil"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
