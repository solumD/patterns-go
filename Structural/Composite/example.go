package composite

import "fmt"

func CompositeExample() {
	homeDirectory := &Directory{Name: "Home", Components: []Component{}}
	file1 := &File{Name: "File 1"}
	file2 := &File{Name: "File 2"}
	file3 := &File{Name: "File 3"}

	homeDirectory.AddComponent(file1)
	homeDirectory.AddComponent(file2)
	homeDirectory.AddComponent(file3)

	projectsDirectory := &Directory{Name: "Projects", Components: []Component{}}
	petProjectFile := &File{Name: "Pet-Project"}

	projectsDirectory.AddComponent(petProjectFile)
	homeDirectory.AddComponent(projectsDirectory)

	fmt.Println("Pet-Project:")
	homeDirectory.Search("Pet-Project")
	fmt.Println("File 2:")
	homeDirectory.Search("File 2")
}
