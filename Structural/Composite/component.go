package composite

// интерфейс, которому должны удовлетворять все компоненты
type Component interface {
	Search(string) bool
}
