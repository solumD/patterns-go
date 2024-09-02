package strategy

type Blog struct {
    strategy Strategy
}

// Устанавливает стратегию конкретной социальной сети 
func (b *Blog) SetStrategy(str Strategy) {
    b.strategy = str
}