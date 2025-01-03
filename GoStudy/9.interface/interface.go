package main

import "fmt"

func myFunc(a any) {
	fmt.Println(a)
}

type AnimalIF interface {
	Sleep()
	GetColor()
	GetType() string
}

type Dog struct {
	color      string
	animalType string
}

func (dog *Dog) GetColor() string {
	return dog.color
}

func (dog *Dog) GetType() string {
	return dog.animalType
}

func main() {
	dog := &Dog{color: "black", animalType: "dog"}
	// 输出: black dog
	fmt.Printf("Color: %s, Type: %s\n", dog.GetColor(), dog.GetType())
	myFunc(dog)
}
