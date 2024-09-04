package bridge

// интерфейс компьютера
type Computer interface {
	LoadGame()
	SetVideocard(Videocard)
}

// интерфейс видеокарты
type Videocard interface {
	Play()
}
