package adapter

import "fmt"

// удовлетворяет интерфейсу Computer
type Mac struct {
}

func (m *Mac) InsertIntoLightning() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}