package mediator

func MediatorExample() {
	dispatcher := NewDispatcher()

	PassengerPlane := &PassengerPlane{
		mediator: dispatcher,
	}
	CargoPlane := &CargoPlane{
		mediator: dispatcher,
	}

	CargoPlane.RequestArrival()
	PassengerPlane.RequestArrival()
	CargoPlane.Depart()
}
