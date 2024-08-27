package factorymethod

import "fmt"

func FactoryMethodExample() {
	products := []Product{}
	c := NewCreator()

	productTypes := []string{phone, "headphones", computer}
	for _, t := range productTypes {
		p, err := c.CreateProduct(t)
		if err != nil {
			fmt.Println(err)
		} else {
			products = append(products, p)
		}
	}

	for _, p := range products {
		p.PrintPrice()
	}
}
