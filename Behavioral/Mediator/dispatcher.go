package mediator

import "sync"

type Dispatcher struct {
	RunwayFree bool
	Queue      []Plane
	lock       *sync.Mutex
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		RunwayFree: true,
		lock:       &sync.Mutex{},
	}
}

// Проверяет, свободна ли полоса, если да, то занимает ее и разрешает прибытие,
// иначе добавляет самолет в очередь
func (d *Dispatcher) CanArrive(p Plane) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.RunwayFree {
		d.RunwayFree = false
		return true
	}
	d.Queue = append(d.Queue, p)
	return false
}

// Уведомляет следующий самолет об отправке, освобождает полосу и осуществляет
// прибытие следующего самолета
func (d *Dispatcher) NotifyAboutDeparture() {
	if !d.RunwayFree {
		d.RunwayFree = true
	}
	if len(d.Queue) > 0 {
		firstPlane := d.Queue[0]
		d.Queue = d.Queue[1:]
		firstPlane.PermitArrival()
	}
}
