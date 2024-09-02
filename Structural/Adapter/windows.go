package adapter

import "fmt"

// не удовлетворяет интерфейсу Computer
type Windows struct{
}

func (w *Windows) insertIntoUSB() {
    fmt.Println("USB connector is plugged into windows machine.")
}