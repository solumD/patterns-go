package bridge

import "fmt"

// домашний компьютер
type HomePC struct {
	videocard Videocard
}

func (c HomePC) LoadGame() {
	fmt.Println("Home PC is loading a game")
	c.videocard.Play()
}

func (c *HomePC) SetVideocard(v Videocard) {
	c.videocard = v
}

// рабочий компьютер
type WorkPC struct {
	videocard Videocard
}

func (c WorkPC) LoadGame() {
	fmt.Println("Work PC is loading a game")
	c.videocard.Play()
}

func (c *WorkPC) SetVideocard(v Videocard) {
	c.videocard = v
}
