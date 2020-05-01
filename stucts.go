package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (p *Person) setName(name string) {
	p.Name = name
}

func main() {
	first := Person{1, "Vasya"}
	fmt.Println(first)
	first.setName("Vasiliy")
	fmt.Printf("%#v \n", first)
}
