package chainofresponsibility

import "fmt"

func ChainOfResponsibilityExample() {
	get := &GetService{Name: "Get 1"}
	validate := &ValidateService{Name: "Validate 1"}
	save := &SaveService{Name: "Save 1"}

	get.SetNext(validate)
	validate.SetNext(save)

	data := &Data{Payload: []byte("some data")}
	err := get.Execute(data)
	if err != nil {
		fmt.Println(err)
	}

	err = get.Execute(data)
	if err != nil {
		fmt.Println(err)
	}
}
