package facade

import "fmt"

// Банк, часть фасада
type Bank struct {
	Name           string
	CardAndBalance map[string]int
}

// Найти карту в банке и вернуть ее баланс
func (b Bank) GetCardBalance(cardNumber string) (int, error) {
	balance, exist := b.CardAndBalance[cardNumber]
	if !exist {
		return -1, fmt.Errorf("[Банк] Неизвестный номер карты: %s", cardNumber)
	}
	fmt.Printf("[Банк] Баланс карты %s: %d.\n", cardNumber, balance)
	return balance, nil
}

// Найти карту в банке и изменить ее баланс
func (b *Bank) ChangeCardBalance(cardNumber string, diff int) error {
	balance, exist := b.CardAndBalance[cardNumber]
	if !exist {
		return fmt.Errorf("[Банк] Неизвестный номер карты: %s", cardNumber)
	}
	if balance+diff < 0 {
		return fmt.Errorf("[Банк] На карте %s недостаточно средств. Баланс: %d", cardNumber, balance)
	}
	b.CardAndBalance[cardNumber] += diff
	fmt.Printf("[Банк] Операция по карте %s успешна. Баланс: %d.\n", cardNumber, b.CardAndBalance[cardNumber])
	return nil
}
