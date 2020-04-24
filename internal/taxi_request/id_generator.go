package taxi_request

import (
	"github.com/valyala/fastrand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func GenerateRequestId() string {
	b := make([]rune, 2)
	for i := range b {
		b[i] = letters[fastrand.Uint32n(uint32(len(letters)))]
	}
	return string(b)
}
