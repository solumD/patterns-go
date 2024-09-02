package adapter

import "fmt"

// Интерфейс компьютера с lighting
type Computer interface {
	InsertIntoLightning()
}

// Клиент, который выполненяет соединение с lighting
type Client struct {
}

// метод имеет какую-то полезную логику, но может взаимодействовать только
// с объектами интерфейса Computer, поэтому при необходимости нужно написать
// адаптер для других объектов
func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightning()
}