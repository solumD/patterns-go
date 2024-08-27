package factorymethod

import "fmt"

const (
	phone    = "phone"
	computer = "computer"
)

type ProductsCreator struct {
}

func NewCreator() Creator {
	return ProductsCreator{}
}

func (c ProductsCreator) CreateProduct(product string) (Product, error) {
	switch product {
	case phone:
		return NewPhone(), nil
	case computer:
		return NewComputer(), nil
	default:
		return nil, fmt.Errorf("%s неизвестный тип", product)
	}
}
