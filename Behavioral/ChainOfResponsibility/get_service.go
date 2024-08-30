package chainofresponsibility

import "fmt"

type GetService struct {
	Name string
	Next Service
}

func (g GetService) Execute(data *Data) error {
	if data.IsGot {
		return fmt.Errorf("data was already got")
	}

	data.IsGot = true
	fmt.Printf("Service %s: data got\n", g.Name)

	err := g.Next.Execute(data)
	if err != nil {
		return err
	}

	return nil
}

func (g *GetService) SetNext(nextService Service) {
	g.Next = nextService
}
