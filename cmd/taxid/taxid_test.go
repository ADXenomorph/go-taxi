package main_test

import (
	"testing"

	main "github.com/ADXenomorph/go-taxi/cmd/taxid"
	"github.com/ADXenomorph/go-taxi/internal/taxi"
	"github.com/ADXenomorph/go-taxi/internal/taxi_request"

	"github.com/stretchr/testify/assert"
)

func TestCreateRouter(t *testing.T) {
	// arrange
	app := taxi.NewApp(taxi_request.NewStorage())

	// act
	router := main.CreateRouter(app)

	// assert
	assert.NotNil(t, router)
}