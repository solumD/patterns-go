package mediator

import "fmt"

type CargoPlane struct {
	mediator PlaneMediator
}

func (p *CargoPlane) Depart() {
	fmt.Println("CargoPlane: отправка")
	p.mediator.NotifyAboutDeparture()
}

func (p *CargoPlane) PermitArrival() {
	fmt.Println("CargoPlane: прибытие разрешено, прибытие")
}

func (p *CargoPlane) RequestArrival() {
	if !p.mediator.CanArrive(p) {
		fmt.Println("CargoPlane: прибытие невозможно, ожидайте")
	} else {
		fmt.Println("CargoPlane: прибытие")
	}
}
