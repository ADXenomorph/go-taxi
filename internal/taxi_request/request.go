package taxi_request

type Status int

const (
	Open Status = 1
	Cancelled Status = 0
)

type Request struct {
	ID     string
	Status Status
}

func NewRequest(id string) *Request {
	return &Request{ID: id, Status: Open}
}
