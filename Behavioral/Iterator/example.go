package iterator

import "fmt"

func IteratorExample() {
	routes := []*Route{
		{"Route 1", 100},
		{"Route 2", 200},
		{"Route 3", 300},
		{"Route 4", 400},
	}

	rIterator := NewRouteIterator(routes)

	fmt.Println("From last to first:")
	rIterator.ToEnd()
	for ; rIterator.HasPrev(); rIterator.PrevIndex() {
		fmt.Println(rIterator.CurrentValue())
	}

	fmt.Println("\nFrom first to last:")
	rIterator.ToStart()
	for ; rIterator.HasNext(); rIterator.NextIndex() {
		fmt.Println(rIterator.CurrentValue())
	}
}
