package task21

import "fmt"

type Dog struct{}

func (dog *Dog) Woof() {
	fmt.Println("Woof, Woof, Woof")
}

type Cat struct{}

func (cat *Cat) Meow() {
	fmt.Println("Meeeoooooowww")
}
