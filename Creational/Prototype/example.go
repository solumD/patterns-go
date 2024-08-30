package prototype

import "fmt"

func PrototypeExample() {
	fmt.Println("Before change:")
	firstStudent := NewStudent("Alex", "Johnson", 19)
	fmt.Printf("Name: %s | Age: %d\n", firstStudent.GetFullName(), firstStudent.GetAge())

	clonedStudent := firstStudent.Clone()
	fmt.Printf("Name: %s | Age: %d\n", clonedStudent.GetFullName(), clonedStudent.GetAge())

	fmt.Println("\nAfter change:")
	clonedStudent.ChangeFullName("Kimmy", "Anderson")
	fmt.Printf("Name: %s | Age: %d\n", firstStudent.GetFullName(), firstStudent.GetAge())
	fmt.Printf("Name: %s | Age: %d\n", clonedStudent.GetFullName(), clonedStudent.GetAge())
}
