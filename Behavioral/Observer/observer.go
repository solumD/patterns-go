package observer

// Интерфейс наблюдателя (подписчика)
type Observer interface {
	Update(string)
	GetName() string
}

// Интерфейс издателя
type Publisher interface {
	Subscribe(Observer)
	UnSubscribe(Observer)
	SetState(string)
	Notify()
}
