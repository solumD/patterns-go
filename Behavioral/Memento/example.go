package memento

import "fmt"

func MementoExample() {
	keeper := &MementoKeeper{
		mementos: make([]*DatabaseMemento, 0),
	}

	originator := &DatabaseOriginator{
		state: "first",
	}

	fmt.Printf("1. Current state: %s\n", originator.GetState())
    keeper.AddMemeto(originator.CreateMemento()) // создаем снимок и добавляет его в хранителя снимков

    originator.SetState("second")
    fmt.Printf("2. Current state: %s\n", originator.GetState())
    keeper.AddMemeto(originator.CreateMemento()) // создаем снимок и добавляет его в хранителя снимков

    originator.SetState("third")
    fmt.Printf("3. Current state: %s\n", originator.GetState())

    originator.RestoreMemento(keeper.GetLatestMemento()) // получаем из хранителя последний снимок и меняем состояние на него 
    fmt.Printf("Restored to state: %s\n", originator.GetState())
 
    originator.RestoreMemento(keeper.GetLatestMemento()) // получаем из хранителя последний снимок и меняем состояние на него 
    fmt.Printf("Restored to state: %s\n", originator.GetState())
}