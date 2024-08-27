package factorymethod

import "fmt"

type Computer struct {
	Name  string
	Price int
}

func NewComputer() Product {
	return Computer{
		Name:  "Macbook", // для наглядного вывода
		Price: 119000,
	}
}

func (p Computer) PrintPrice() {
	fmt.Printf("Phone %s, price %d\n", p.Name, p.Price)
}
