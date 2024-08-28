package builder

import "fmt"

// Интерфейс строителя, которому должны удовлетворять
// строители каждого бренда гитары
type Builder interface {
	SetBrand()
	SetStringsQuantity()
	SetColour()
	GetGuitar() Guitar
}

const (
	FenderGuitar = "fender"
	IbanezGuitar = "ibanez"
)

// возвращает строителя конкретной гитары
func GetBuilder(guitarBrand string) (Builder, error) {
	switch guitarBrand {
	default:
		return nil, fmt.Errorf("invalid brand: %s", guitarBrand)
	case FenderGuitar:
		return &FenderBuilder{}, nil
	case IbanezGuitar:
		return &IbanezBuilder{}, nil
	}
}
