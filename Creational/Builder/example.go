package builder

import "fmt"

func BuilderExample() {
	fenderBuilder, err := GetBuilder("fender")
	if err != nil {
		fmt.Println(err)
	}

	guitarFactory := NewGuitarBuilder(fenderBuilder)

	fenderGuitar := guitarFactory.CreateGuitar()
	fenderGuitar.PrintDetails()

	ibanezBuilder, _ := GetBuilder("ibanez")
	if err != nil {
		fmt.Println(err)
	}
	guitarFactory.SetBuilder(ibanezBuilder)

	ibanezGuitar := guitarFactory.CreateGuitar()
	ibanezGuitar.PrintDetails()
}
