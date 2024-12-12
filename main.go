package main

import "fmt"

type Human struct {
	age    int
	height int
	weigh  int
}

func (h Human) HumanPrint() {
	fmt.Printf("age = %d height = %d weigh = %d", h.age, h.height, h.weigh)
}

type Action struct {
	Human
	x int
}

func main() {
	act := Action{
		Human: Human{age: 10,
			height: 15,
			weigh:  20,
		},
		x: 3,
	}
	act.HumanPrint()
}
