package decorator

// Карлсон с пропеллером
type CarlsonWithPropeller struct {
	Carlson CarlsonDecorator
}

func (pc CarlsonWithPropeller) GetSpeed() int {
	return pc.Carlson.GetSpeed() + 5
}
