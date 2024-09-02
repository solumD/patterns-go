package adapter

import "fmt"

// адаптирует тип windows, чтобы он мог работать с клиентом
type WindowsAdapter struct {
	windowsComputer *Windows
}

func (a *WindowsAdapter) InsertIntoLightning() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	a.windowsComputer.insertIntoUSB()
}