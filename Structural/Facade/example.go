package facade

import "fmt"

func FacadeExample() {
	shop := NewGuitarShopFacade()
	bank := &Bank{Name: "Tinkoff", CardAndBalance: map[string]int{"1111222233334444": 100000}}
	card := &Card{Number: "1111222233334444", Bank: bank}

	fmt.Println("Успешная операция.")
	shop.SellGuitar("Fender", card)
	fmt.Println("\nОперация с недостаточным кол-вом средств.")
	shop.SellGuitar("Ibanez", card)
}
