package observer

type NewsPublisher struct {
	State     string
	Observers map[string]Observer
}

func NewNewsPublisher() Publisher {
	return &NewsPublisher{
		State:     "basic state",
		Observers: map[string]Observer{},
	}
}

func (p *NewsPublisher) Subscribe(o Observer) {
	p.Observers[o.GetName()] = o
}

func (p *NewsPublisher) UnSubscribe(o Observer) {
	delete(p.Observers, o.GetName())
}

func (p *NewsPublisher) SetState(s string) {
	p.State = s
}

func (p *NewsPublisher) Notify() {
	for _, o := range p.Observers {
		o.Update(p.State)
	}
}
