package composite

import "fmt"

// структура "директория", у которой есть "дети", удовлетворяющие интерфейсу Component
type Directory struct {
	Name       string
	Components []Component
}

func (d Directory) Search(name string) bool {
	if d.Name == name {
		fmt.Printf("Component with name %s found\n", name)
		return true
	}

	fmt.Printf("Searching in the directory %s\n", d.Name)
	for _, c := range d.Components {
		if found := c.Search(name); found {
			return true
		}
	}
	return false
}

func (d *Directory) AddComponent(cmp Component) {
	d.Components = append(d.Components, cmp)
}
