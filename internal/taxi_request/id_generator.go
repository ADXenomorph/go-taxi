package taxi_request

import (
	"github.com/valyala/fastrand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

// GenerateRequestId creates a random string of 2 latin lowercase letters
func GenerateRequestId() string {
	b := make([]rune, 2)
	for i := range b {
		// fastrand usage allowed to reduce operation time from 8000ns to 80ns
		b[i] = letters[fastrand.Uint32n(uint32(len(letters)))]
	}
	return string(b)
}
