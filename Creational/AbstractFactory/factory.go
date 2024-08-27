package abstractfactory

type Factory interface {
	GetApp(int, int) App
	GetStaff(int) Staff
}

type App interface {
	PrintUsers()
	PrintGain()
}

type Staff interface {
	PrintQuantity()
}
