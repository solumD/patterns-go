package abstractfactory

import "fmt"

type OzonFactory struct {
}

func NewOzonFactory() OzonFactory {
	return OzonFactory{}
}

func (o OzonFactory) GetApp(u int, g int) App {
	return OzonApp{
		Users: u,
		Gain:  g,
	}
}

func (o OzonFactory) GetStaff(q int) Staff {
	return OzonStaff{
		Quantity: q,
	}
}

type OzonApp struct {
	Users int
	Gain  int
}

func (a OzonApp) PrintUsers() {
	fmt.Printf("[Ozon App] Users: %d\n", a.Users)
}

func (a OzonApp) PrintGain() {
	fmt.Printf("[Ozon App] Gain: %d\n", a.Gain)
}

type OzonStaff struct {
	Quantity int
}

func (s OzonStaff) PrintQuantity() {
	fmt.Printf("[Ozon Staff] Quantity: %d\n", s.Quantity)
}
