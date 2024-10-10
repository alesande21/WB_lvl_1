package main

import "fmt"

type Human struct {
	name   string
	age    int
	height int
	weight float32
}

func (h *Human) Run() {
	fmt.Printf("Человек %s бежит!\n", h.name)
}

func (h *Human) Walk() {
	fmt.Printf("Человек %s идет!\n", h.name)
}

func (h *Human) Sleep() {
	fmt.Printf("Человек %s спит!\n", h.name)
}

func (h *Human) Eat() {
	fmt.Printf("Человек %s ест! Рост/вес: %d/%f\n", h.name, h.height, h.weight)
}

type Action struct {
	Human
}

func getAction(name string, age int, height int, weight float32) *Action {
	return &Action{struct {
		name   string
		age    int
		height int
		weight float32
	}{name: name, age: age, height: height, weight: weight}}
}

func main() {
	act := getAction("Антон", 74, 174, 74)
	act.Walk()
	act.Sleep()
	act.Run()
	act.Eat()
}
