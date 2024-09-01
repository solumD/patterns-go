package mediator

// Посредник всех самолетов
type PlaneMediator interface {
	CanArrive(Plane) bool
	NotifyAboutDeparture()
}

// Интерфейс для всех самолетов
type Plane interface {
	Depart()
	PermitArrival()
	RequestArrival()
}
