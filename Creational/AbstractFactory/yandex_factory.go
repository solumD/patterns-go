package abstractfactory

import "fmt"

type YandexFactory struct {
}

func NewYandexFactory() YandexFactory {
	return YandexFactory{}
}

func (o YandexFactory) GetApp(u int, g int) App {
	return YandexApp{
		Users: u,
		Gain:  g,
	}
}

func (o YandexFactory) GetStaff(q int) Staff {
	return YandexStaff{
		Quantity: q,
	}
}

type YandexApp struct {
	Users int
	Gain  int
}

func (a YandexApp) PrintUsers() {
	fmt.Printf("[Yandex App] Users: %d\n", a.Users)
}

func (a YandexApp) PrintGain() {
	fmt.Printf("[Yandex App] Gain: %d\n", a.Gain)
}

type YandexStaff struct {
	Quantity int
}

func (s YandexStaff) PrintQuantity() {
	fmt.Printf("[Yandex Staff] Quantity: %d\n", s.Quantity)
}
