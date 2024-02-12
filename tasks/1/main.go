package main

import "fmt"

/*
	Дана структура Human (с произвольным набором полей и методов). 
	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

// Структура Human с полями, представляющими человека
type Human struct {
	lastName  string
	firstName string
	age       uint8
	growth    uint8
	gender    string
}

// Action - структура, встраивающая (наследующая) Human
type Action struct {
	// Встраивание Human в Action позволяет использовать его поля и методы
	Human
	count int
}

// Метод для Human
func (h Human) introduceYourself() {
	fmt.Printf("Hi, my name is %s %s. I'm %d years old. I am %s and my height is %d cm\n",
		h.firstName, h.lastName, h.age, h.gender, h.growth)
}

// Метод для Action
func (a Action) sayHello() {
	/*
		В следующей строчке можно заметить, что мы явно не указываем структуру Human
		для получения lastName и firstName, так как мы встроили (унаследовали) эти поля
		от родительской структуры Human. Однако мы можем явно указать a.Human.firstName,
		это будет необходимо, если в структуре Action есть поля с таким же именем,
		и мы хотим обратиться конкретно к a.Human.firstName.
	*/
	myName := fmt.Sprintf("%s %s", a.firstName, a.lastName)
	for i := 0; i < a.count; i++ {
		fmt.Printf("Hi, my name is %s.\n", myName)
	}
}

func main() {
	act := Action{
		Human: Human{
			firstName: "Artem",
			lastName:  "Echeistov",
			age:       25,
			growth:    170,
			gender:    "male",
		},
		count: 2,
	}
	/*
		Мы можем не указывать явно Human при вызове метода introduceYourself(),
		так как структура Human встроена в структуру Action (Action унаследовала методы от Human).
		В языке Go нам не нужно явно указывать, что Action наследует методы от Human,
		это делается благодаря встраиванию структуры Human в Action.
	*/
	act.introduceYourself()
	act.sayHello()
}
