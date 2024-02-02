package task1

import "fmt"

type Human struct {
	Name, Sex string
	Age       int
}

type Action struct {
	Human
}

func (h *Human) printInfo() {

	fmt.Printf("Name: %s\nAge: %d\nSex: %s", h.Name, h.Age, h.Sex)
}

func Run() {
	a := Action{Human{Name: "Ivan", Sex: "sometimes"}}
	a.printInfo()
}
