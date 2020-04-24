package taxi_test

import (
	"testing"

	"github.com/ADXenomorph/go-taxi/internal/taxi"
	"github.com/ADXenomorph/go-taxi/internal/taxi_request"

	"github.com/stretchr/testify/assert"
)

func TestCreateRequest(t *testing.T) {
	// arrange
	app := taxi.NewApp(taxi_request.NewStorage())

	// act
	req := app.CreateRequest("aa")

	// assert
	assert.Equal(t, "aa", req.ID)
	assert.Equal(t, taxi_request.Open, req.Status)
}

func TestCancelExistingRequest(t *testing.T) {
	// arrange
	app := taxi.NewApp(taxi_request.NewStorage())
	app.CreateRequest("aa")

	// act
	err := app.CancelRequest("aa")

	// assert
	assert.Nil(t, err)
}

func TestCancelMissingRequest(t *testing.T) {
	// arrange
	app := taxi.NewApp(taxi_request.NewStorage())

	// act
	err := app.CancelRequest("aa")

	// assert
	assert.NotNil(t, err)
	assert.EqualError(t, err, "Failed to cancel")
}

func TestGetRandomRequestWithEmptySet(t *testing.T) {
	// arrange
	app := taxi.NewApp(taxi_request.NewStorage())

	// act
	req := app.GetRandomRequest()

	// assert
	assert.Nil(t, req)
}

func TestGetRequestStatisticsEmpty(t *testing.T) {
	// arrange
	app := taxi.NewApp(taxi_request.NewStorage())

	// act
	res := app.GetRequestStatistics()

	// assert
	assert.NotNil(t, res)
	assert.Len(t, res, 0)
}

func TestGetRequestStatisticsNotEmpty(t *testing.T) {
	// arrange
	app := taxi.NewApp(taxi_request.NewStorage())
	app.CreateRandomRequest()
	app.GetRandomRequest()

	// act
	res := app.GetRequestStatistics()

	// assert
	assert.NotNil(t, res)
	assert.Len(t, res, 1)
}

func TestCreateRandomRequest(t *testing.T) {
	// arrange
	app := taxi.NewApp(taxi_request.NewStorage())

	// act
	req := app.CreateRandomRequest()

	// assert
	assert.NotNil(t, req)
	assert.Equal(t, taxi_request.Open, req.Status)
}

func TestCreateInitialRequests(t *testing.T) {
	// arrange
	app := taxi.NewApp(taxi_request.NewStorage())

	// act
	app.CreateInitialRequests()
}

func TestCancelRandomRequestWithEmptySet(t *testing.T) {
	// arrange
	app := taxi.NewApp(taxi_request.NewStorage())

	// act
	app.CancelRandomRequest()
}

func TestCancelRandomRequest(t *testing.T) {
	// arrange
	storage := taxi_request.NewStorage()
	app := taxi.NewApp(storage)
	app.CreateRequest("aa")

	// act
	app.CancelRandomRequest()
	req, ok := storage.Get("aa")

	// assert
	assert.True(t, ok)
	assert.NotNil(t, req)
	assert.Equal(t, taxi_request.Cancelled, req.Status)
}
