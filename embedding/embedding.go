// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

// Human структура с произвольным набором полей и методов
type Human struct {
	Name string
	Age  int
}

// Метод Human
func (h *Human) Speak() {
	fmt.Printf("Hi, my name is %s and I am %d years old.\n", h.Name, h.Age)
}

// Action структура, включающая Human
type Action struct {
	Human
	Profession string
}

// Метод Action
func (a *Action) Work() {
	fmt.Printf("%s is working as a %s.\n", a.Name, a.Profession) // использование поля из встроенной структуры
}

func main() {
	// Создание экземпляра Action
	a := Action{
		Human:      Human{Name: "John", Age: 30},
		Profession: "Engineer",
	}

	// Вызов методов Human через экземпляр Action
	a.Speak()
	a.Work()
}
