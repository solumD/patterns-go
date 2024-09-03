package proxy

import "fmt"

func ProxyExample() {
	proxy := &GoodsServiceProxy{
		realService: &GoodsService{},
	}

	fmt.Println("Invalid case")
	goods, err := proxy.GetAllFruits("some invalid key")
	if err != nil {
		fmt.Println(goods, err)
	}

	fmt.Println("\nValid case")
	goods, _ = proxy.GetAllFruits("go is awesome")
	fmt.Println(goods)
}
