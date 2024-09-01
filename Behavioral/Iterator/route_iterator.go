package iterator

type Route struct {
	Name   string
	Metres int
}

type RouteIterator struct {
	Routes []*Route
	Index  int
}

func NewRouteIterator(r []*Route) Iterator {
	return &RouteIterator{
		Routes: r,
		Index:  0,
	}
}

func (r RouteIterator) HasNext() bool {
	return r.Index < len(r.Routes)
}

func (r RouteIterator) HasPrev() bool {
	return r.Index >= 0
}

func (r *RouteIterator) NextIndex() {
	if r.HasNext() {
		r.Index++
	}
}

func (r *RouteIterator) PrevIndex() {
	if r.HasPrev() {
		r.Index--
	}
}

func (r RouteIterator) CurrentValue() interface{} {
	return r.Routes[r.Index]
}

func (r *RouteIterator) ToStart() {
	r.Index = 0
}

func (r *RouteIterator) ToEnd() {
	r.Index = len(r.Routes) - 1
}
