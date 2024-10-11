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
	name  string // отдельно определили поле для структуры Action
	Human        // встраивание структуры Human. Методы и поля Human автоматически доступны в Action.
}

func getAction(nameAction string, nameHuman string, age int, height int, weight float32) *Action {
	return &Action{
		name: nameAction,
		Human: struct {
			name   string
			age    int
			height int
			weight float32
		}{name: nameHuman, age: age, height: height, weight: weight},
	}
}

func main() {
	act := getAction("режим спорт", "Антон", 74, 174, 74)
	act.Walk()
	act.Sleep()
	act.Run()
	act.Eat()
	/*
		Если имена полей совпадают в структуре Human и Action, приоритет будет отдан полю основной структуры Action.
		Чтобы обратиться к полю родительской структуры Human, нужно явно указать структуру, например:
			nameAction := act.name       // Обращение к полю name структуры Action.
			nameHuman := act.Human.name  // Обращение к полю name структуры Human.
	*/
}
