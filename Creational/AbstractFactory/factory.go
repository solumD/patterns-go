package abstractfactory

// Общий интерфейс фабрики
type Factory interface {
	GetApp(int, int) App
	GetStaff(int) Staff
}

// Общий интерфейса продукта А
type App interface {
	PrintUsers()
	PrintGain()
}

// Общий интерфейса продукта Б
type Staff interface {
	PrintQuantity()
}
