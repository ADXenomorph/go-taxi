package taxi_request_test

import (
	"testing"

	"github.com/ADXenomorph/go-taxi/internal/taxi_request"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	// arrange
	s := taxi_request.NewStorage()

	// act
	s.Save(taxi_request.NewRequest("aa"))
}

func TestGetWithoutRequest(t *testing.T) {
	// arrange
	s := taxi_request.NewStorage()

	// act
	req, ok := s.Get("aa")

	// assert
	assert.Nil(t, req)
	assert.False(t, ok)
}

func TestGetWithExistingRequest(t *testing.T) {
	// arrange
	s := taxi_request.NewStorage()
	s.Save(taxi_request.NewRequest("aa"))

	// act
	req, ok := s.Get("aa")

	// assert
	assert.NotNil(t, req)
	assert.True(t, ok)
}

func TestGetRandomForEmptySet(t *testing.T) {
	// arrange
	s := taxi_request.NewStorage()

	// act
	req := s.GetRandom()

	// assert
	assert.Nil(t, req)
}

func TestGetRandom(t *testing.T) {
	s := taxi_request.NewStorage()
	// arrange
	s.Save(taxi_request.NewRequest("aa"))

	// act
	req := s.GetRandom()

	// assert
	assert.NotNil(t, req)
}

func TestGetRandomAndCountWithEmptySet(t *testing.T) {
	// arrange
	s := taxi_request.NewStorage()

	// act
	req := s.GetRandomAndCount()

	// assert
	assert.Nil(t, req)
}

func TestGetRandomAndCount(t *testing.T) {
	// arrange
	s := taxi_request.NewStorage()
	s.Save(taxi_request.NewRequest("aa"))

	// act
	req := s.GetRandomAndCount()

	// assert
	assert.NotNil(t, req)
}

func TestGetCountersWithEmptySet(t *testing.T) {
	// arrange
	s := taxi_request.NewStorage()

	// act
	res := s.GetCounters()

	// assert
	assert.Equal(t, res, []string{})
}

func TestGetCountersWithRequestAndNoCounters(t *testing.T) {
	// arrange
	s := taxi_request.NewStorage()
	s.Save(taxi_request.NewRequest("aa"))

	// act
	res := s.GetCounters()

	// assert
	assert.NotNil(t, res)
	assert.Len(t, res, 0)
}

func TestGetCountersWithRequestAndCounters(t *testing.T) {
	// arrange
	s := taxi_request.NewStorage()
	s.Save(taxi_request.NewRequest("aa"))
	s.GetRandomAndCount()

	// act
	res := s.GetCounters()

	// assert
	assert.NotNil(t, res)
	assert.Len(t, res, 1)
}
