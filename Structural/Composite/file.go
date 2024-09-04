package composite

import "fmt"

// структура "файл", удовлетворяющая интерфейсу Component
type File struct {
	Name string
}

func (f File) Search(name string) bool {
	if f.Name == name {
		fmt.Printf("Component with name %s found\n", name)
		return true
	}
	return false
}
