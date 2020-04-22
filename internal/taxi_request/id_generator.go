package taxi_request

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func GenerateRequestId() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 2)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
