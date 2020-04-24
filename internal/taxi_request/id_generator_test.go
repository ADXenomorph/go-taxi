package taxi_request_test

import (
	"testing"

	"github.com/ADXenomorph/go-taxi/internal/taxi_request"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRequestId(t *testing.T) {
	// act
	id := taxi_request.GenerateRequestId()

	// assert
	assert.NotNil(t, id)
	assert.Len(t, id, 2)
}

func BenchmarkGenerateRequestId(b *testing.B) {
	for i := 0; i < b.N; i++ {
        taxi_request.GenerateRequestId()
    }
}