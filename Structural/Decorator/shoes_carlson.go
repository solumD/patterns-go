package decorator

// Карлсон в обуви
type CarlsonWithShoes struct {
	Carlson CarlsonDecorator
}

func (sc CarlsonWithShoes) GetSpeed() int {
	return sc.Carlson.GetSpeed() + 1
}
