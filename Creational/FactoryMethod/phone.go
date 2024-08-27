package factorymethod

import "fmt"

type Phone struct {
	Name  string
	Price int
}

func NewPhone() Phone {
	return Phone{
		Name:  "IPhone 12", // для наглядного вывода
		Price: 67000,
	}
}

func (p Phone) PrintPrice() {
	fmt.Printf("Phone %s, price %d\n", p.Name, p.Price)
}
