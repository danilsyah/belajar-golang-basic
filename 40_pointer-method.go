package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	danil := Man{"Danil"}
	danil.Married()

	fmt.Println(danil.Name)
}
