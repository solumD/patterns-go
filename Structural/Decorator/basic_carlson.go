package decorator

// Карлсон без всего
type BasicCarlson struct {
}

func (bc BasicCarlson) GetSpeed() int {
	return 0
}
