package builder

// "Директор" всех строителей, который ими управляет
type GuitarBuilder struct {
	Builder Builder
}

// Создать директора для управления переданным строителем
func NewGuitarBuilder(b Builder) *GuitarBuilder {
	return &GuitarBuilder{Builder: b}
}

// Заменить строителя текущего бренда, на строителя другого бренда
func (gb *GuitarBuilder) SetBuilder(b Builder) {
	gb.Builder = b
}

// Дает команду строителю создать гитару, а затем вернуть ее
func (gb *GuitarBuilder) CreateGuitar() Guitar {
	gb.Builder.SetBrand()
	gb.Builder.SetStringsQuantity()
	gb.Builder.SetColour()
	return gb.Builder.GetGuitar()
}
