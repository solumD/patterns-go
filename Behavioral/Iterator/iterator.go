package iterator

type Iterator interface {
	HasNext() bool
	HasPrev() bool
	NextIndex()
	PrevIndex()
	CurrentValue() interface{}
	ToStart()
	ToEnd()
}
