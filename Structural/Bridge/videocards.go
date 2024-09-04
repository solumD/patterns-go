package bridge

import "fmt"

type RTX2080 struct {
}

func (v RTX2080) Play() {
	fmt.Println("Playing game in 300 fps")
}

type GTX1060 struct {
}

func (v GTX1060) Play() {
	fmt.Println("Playing game in 150 fps")
}
