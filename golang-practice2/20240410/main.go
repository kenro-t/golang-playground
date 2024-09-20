package main

import "fmt"

func main() {
	dog := Dog{name: "taro"}
	cat := Cat{name: "tama"}

	animals := []Animal{dog, cat}
	for _, animal := range animals {
		fmt.Println(animal.GetName(), ":", animal.Speak())
	}
}

type Animal interface {
	Speak() string
	GetName() string
}

type Dog struct {
	name string
}
type Cat struct {
	name string
}

func (d Dog) Speak() string {
	return "Bow"
}

func (c Cat) Speak() string {
	return "Meow"
}

func (d Dog) GetName() string {
	return d.name
}

func (c Cat) GetName() string {
	return c.name
}
