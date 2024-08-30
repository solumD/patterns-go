package chainofresponsibility

import "fmt"

type SaveService struct {
	Name string
	Next Service
}

func (s SaveService) Execute(data *Data) error {
	if data.IsSaved {
		return fmt.Errorf("data was already saved")
	}

	data.IsSaved = true
	fmt.Printf("Service %s: data saved\n", s.Name)

	return nil
}

func (s *SaveService) SetNext(nextService Service) {
	s.Next = nextService
}
