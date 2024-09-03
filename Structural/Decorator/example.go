package decorator

import "fmt"

func DecoratorExample() {

	bc := &BasicCarlson{}
	sc := &CarlsonWithShoes{
		Carlson: bc,
	}
	pc := &CarlsonWithPropeller{
		Carlson: bc,
	}

	fmt.Printf("Basic Carlson's speed: %d\n", bc.GetSpeed())
	fmt.Printf("Carlson's with shoes speed: %d\n", sc.GetSpeed())
	fmt.Printf("Carlson's with propeller speed: %d\n", pc.GetSpeed())
}
