package builder

// Строитель гитары бренда Ibanez,
// создает базовую комплектацию
type IbanezBuilder struct {
	Brand           string
	StringsQuantity int
	Colour          string
}

func (b *IbanezBuilder) SetBrand() {
	b.Brand = "ibanez"
}

func (b *IbanezBuilder) SetStringsQuantity() {
	b.StringsQuantity = 7
}

func (b *IbanezBuilder) SetColour() {
	b.Colour = "black"
}

func (b *IbanezBuilder) GetGuitar() Guitar {
	return Guitar{
		Brand:           b.Brand,
		StringsQuantity: b.StringsQuantity,
		Colour:          b.Colour,
	}
}
