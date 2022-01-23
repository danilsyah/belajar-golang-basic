package main

import "fmt"

type HasName interface {
	GetName() string
	GetUmur() int
	GetWarna() string
}

func SayHello(hasName HasName) {
	fmt.Println("Helo", hasName.GetName(), "Umur ", hasName.GetUmur(), "warna", hasName.GetWarna())
}

type Person struct {
	Name  string
	Umur  int
	Warna string
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetUmur() int {
	return person.Umur
}

func (person Person) GetWarna() string {
	return person.Warna
}

type Animal struct {
	Name  string
	Warna string
	Umur  int
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (animal Animal) GetWarna() string {
	return animal.Warna
}

func (animal Animal) GetUmur() int {
	return animal.Umur
}

func main() {
	var danil Person
	danil.Name = "danil syah"
	danil.Umur = 27
	danil.Warna = "Putih"

	SayHello(danil)

	kucing := Animal{
		Name:  "Persia",
		Warna: "Coklat",
		Umur:  1,
	}

	SayHello(kucing)
}
