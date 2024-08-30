package chainofresponsibility

import "fmt"

type ValidateService struct {
	Name string
	Next Service
}

func (v ValidateService) Execute(data *Data) error {
	if data.IsValidated {
		return fmt.Errorf("data was already validated")
	}

	data.IsValidated = true
	fmt.Printf("Service %s: data validated\n", v.Name)

	err := v.Next.Execute(data)
	if err != nil {
		return err
	}

	return nil
}

func (v *ValidateService) SetNext(nextService Service) {
	v.Next = nextService
}
