package mediator

import "fmt"

type PassengerPlane struct {
	mediator PlaneMediator
}

func (p *PassengerPlane) Depart() {
	fmt.Println("PassengerPlane: отправка")
	p.mediator.NotifyAboutDeparture()
}

func (p *PassengerPlane) PermitArrival() {
	fmt.Println("PassengerPlane: прибытие разрешено, посадка")
}

func (p *PassengerPlane) RequestArrival() {
	if !p.mediator.CanArrive(p) {
		fmt.Println("PassengerPlane: прибытие невозможно, ожидайте")
	} else {
		fmt.Println("PassengerPlane: прибытие")
	}
}
