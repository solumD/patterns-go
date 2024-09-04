package bridge

import "fmt"

func BridgeExample() {

	// видеокарты
	rtx2080 := &RTX2080{}
	gtx1060 := &GTX1060{}

	// компьютеры
	homePC := &HomePC{}
	homePC.SetVideocard(rtx2080)
	workPC := &WorkPC{}
	workPC.SetVideocard(gtx1060)

	fmt.Println("Home pc - rtx 2080, Work pc - gtx 1060")
	fmt.Println()
	homePC.LoadGame()
	workPC.LoadGame()

	homePC.SetVideocard(gtx1060)
	workPC.SetVideocard(rtx2080)

	fmt.Println("-----------------------")
	fmt.Println("Home pc - gtx 1060, Work pc - rtx 2080")
	fmt.Println()
	homePC.LoadGame()
	workPC.LoadGame()
}
