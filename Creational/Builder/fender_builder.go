package builder

// Строитель гитары бренда Fender,
// создает базовую комплектацию
type FenderBuilder struct {
	Brand           string
	StringsQuantity int
	Colour          string
}

func (b *FenderBuilder) SetBrand() {
	b.Brand = "fender"
}

func (b *FenderBuilder) SetStringsQuantity() {
	b.StringsQuantity = 6
}

func (b *FenderBuilder) SetColour() {
	b.Colour = "white"
}

func (b *FenderBuilder) GetGuitar() Guitar {
	return Guitar{
		Brand:           b.Brand,
		StringsQuantity: b.StringsQuantity,
		Colour:          b.Colour,
	}
}
