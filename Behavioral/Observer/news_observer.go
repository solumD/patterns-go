package observer

import "fmt"

type NewsObserver struct {
	Name  string
	State string
}

func (o *NewsObserver) Update(s string) {
	fmt.Printf("Observer %s: changing state from [%s] to [%s]\n", o.Name, o.State, s)
	o.State = s
}

func (o *NewsObserver) GetName() string {
	return o.Name
}
