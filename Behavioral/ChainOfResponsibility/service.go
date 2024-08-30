package chainofresponsibility

type Service interface {
	Execute(*Data) error
	SetNext(Service)
}

type Data struct {
	Payload     []byte
	IsGot       bool
	IsValidated bool
	IsSaved     bool
}
