package facade

import "fmt"

// Главный фасад, который управляет дальнешими операциями
type GuitarShopFacade struct {
	Guitars map[string]int
}

func NewGuitarShopFacade() *GuitarShopFacade {
	return &GuitarShopFacade{
		Guitars: map[string]int{"Ibanez": 50000, "Fender": 70000},
	}
}

// Операция продажи гитары по указанному имени гитары и номеру карты
func (s *GuitarShopFacade) SellGuitar(guitarName string, card *Card) {
	fmt.Printf("[Магазин] Поиск гитары с названием %s.\n", guitarName)
	guitarPrice, exist := s.Guitars[guitarName]
	if !exist {
		fmt.Printf("[Магазин] Гитара с названием %s не найдена.\n", guitarName)
		return
	}
	fmt.Printf("[Магазин] Гитара найдена.\n")
	fmt.Printf("[Магазин] Проверка баланса карты: %s.\n", card.Number)
	_, err := card.GetBalance()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[Магазин] Попытка списать средства за покупку гитары.\n")
	err = card.ChangeBalance(-guitarPrice)
	if err != nil {
		fmt.Printf("[Магазин] Ошибка при списывании средств.\n")
		fmt.Println(err)
		return
	}
	fmt.Printf("[Магазин] Покупка гитары завершена.\n")
}
