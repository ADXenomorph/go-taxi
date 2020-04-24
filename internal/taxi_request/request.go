package taxi_request

// Status represents taxi request Open/Cancelled status
type Status int

const (
	Open      Status = 1
	Cancelled Status = 0
)

// Request represents a taxi request
type Request struct {
	ID     string
	Status Status
}

// NewRequest returns new initialized taxi Request
func NewRequest(id string) *Request {
	return &Request{ID: id, Status: Open}
}
