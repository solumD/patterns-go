package builder

import "fmt"

// Объект, который производят все строители
type Guitar struct {
	Brand           string
	StringsQuantity int
	Colour          string
}

func (g Guitar) PrintDetails() {
	fmt.Printf("Brand: %s, Strings Quantity: %d, Colour: %s\n", g.Brand, g.StringsQuantity, g.Colour)
}
