package facade

import "fmt"

// Карта, часть фасада
type Card struct {
	Number string
	Bank   *Bank
}

// Получить баланс карты
func (c Card) GetBalance() (int, error) {
	fmt.Printf("[Карта] Номер %s. Проверка баланса в банке.\n", c.Number)
	balance, err := c.Bank.GetCardBalance(c.Number)
	if err != nil {
		return -1, err
	}
	return balance, nil
}

// Изменить баланс карты
func (c *Card) ChangeBalance(diff int) error {
	fmt.Printf("[Карта] Номер %s. Попытка провести операцию в банке.\n", c.Number)
	err := c.Bank.ChangeCardBalance(c.Number, diff)
	if err != nil {
		return err
	}
	return nil
}
