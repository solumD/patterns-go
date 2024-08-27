package factorymethod

type Creator interface {
	CreateProduct(product string) (Product, error)
}

type Product interface {
	PrintPrice()
}
