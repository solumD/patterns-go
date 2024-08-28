package singleton

import (
	"fmt"
	"sync"
)

var (
	s    *Singleton
	once sync.Once
)

type Singleton struct {
	Name string
}

func InitSingleton(name string) *Singleton {
	once.Do(func() {
		s = &Singleton{Name: name}
	})
	return s
}

func (s Singleton) PrintName() {
	fmt.Printf("Name: %s\n", s.Name)
}
