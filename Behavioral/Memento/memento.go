package memento

// Снимок БД
type DatabaseMemento struct {
    state string 
}

func (m *DatabaseMemento) GetState() string {
    return m.state
}


// Создатель снимков БД
type DatabaseOriginator struct {
    state string
}

func (o *DatabaseOriginator) CreateMemento() *DatabaseMemento {
    return &DatabaseMemento{state: o.state}
}

func (o *DatabaseOriginator) RestoreMemento(m *DatabaseMemento)  {
    o.state = m.GetState()
}

func (o *DatabaseOriginator) SetState(state string) {
    o.state = state
}

func (o *DatabaseOriginator) GetState() string {
    return o.state
}


// Хранитель снимков БД
type MementoKeeper struct {
    mementos []*DatabaseMemento
}

func (k *MementoKeeper) AddMemeto(m *DatabaseMemento) {
    k.mementos = append(k.mementos, m)
}

func(k *MementoKeeper) GetLatestMemento() *DatabaseMemento {
    n := len(k.mementos)
    m := k.mementos[n - 1]
    k.mementos = k.mementos[:n-1]
    return m
}