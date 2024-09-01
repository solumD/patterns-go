package observer

import "fmt"

func ObserverExample() {
	publisher := NewNewsPublisher()

	o1 := &NewsObserver{Name: "observer 1"}
	o2 := &NewsObserver{Name: "observer 2"}
	o3 := &NewsObserver{Name: "observer 3"}
	publisher.Subscribe(o1)
	publisher.Subscribe(o2)
	publisher.Subscribe(o3)

	fmt.Println("Before unsubscribe:")
	publisher.Notify()

	publisher.SetState("new state")
	publisher.UnSubscribe(o2)

	fmt.Println("\nAfter unsubscribe")
	publisher.Notify()
}
