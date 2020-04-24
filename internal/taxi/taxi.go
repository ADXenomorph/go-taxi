/*
Package taxi implements main business logic of the taxi app.

Create the initial state of the app with CreateInitialRequests.
Simulate incoming changes with SimulateChanges.
Start getting requests with GetRandomRequest.
And get request statistics with GetRequestStatistics.
*/
package taxi

import (
	"errors"
	"time"

	"github.com/ADXenomorph/go-taxi/internal/taxi_request"
)

// App represents the taxi app and is the main entry
// point for all the actions that can be done
type App struct {
	storage *taxi_request.RequestStorage
}

// NewApp returns new initialized taxi App
func NewApp(storage *taxi_request.RequestStorage) *App {
	return &App{storage: storage}
}

// CreateRequest creates a taxi request by id, saves it to storage
// and returns it
func (app *App) CreateRequest(id string) *taxi_request.Request {
	tr := taxi_request.NewRequest(id)
	app.storage.Save(tr)

	return tr
}

// CancelRequest loads taxi request by id, sets it's status to Cancelled
// and saves it to storage
func (app *App) CancelRequest(id string) error {
	tr, ok := app.storage.Get(id)

	if !ok {
		return errors.New("Failed to cancel")
	}

	tr.Status = taxi_request.Cancelled
	app.storage.Save(tr)

	return nil
}

// GetRandomRequest returns a random open taxi request and increases
// the statistic counter of this request. It is the main func of the application.
func (app *App) GetRandomRequest() *taxi_request.Request {
	return app.storage.GetRandomAndCount()
}

// GetRequestStatistics returns a slice of statistic lines for all requests
// Requests with zero result are skipped.
// Example: ["aa - 456", "bb - 123"]
func (app *App) GetRequestStatistics() []string {
	return app.storage.GetCounters()
}

// CreateRandomRequest generates a random request id,
// creates a taxi request using the id and returns the request
func (app *App) CreateRandomRequest() *taxi_request.Request {
	return app.CreateRequest(taxi_request.GenerateRequestId())
}

// CreateInitialRequests creates 50 random requests for
// the initial state of the app
func (app *App) CreateInitialRequests() {
	for i := 0; i < 50; i++ {
		app.CreateRandomRequest()
	}
}

// CancelRandomRequest finds a random request and if its found cancels it
func (app *App) CancelRandomRequest() {
	req := app.storage.GetRandom()

	if req == nil {
		return
	}

	app.CancelRequest(req.ID)
}

// SimulateChanges is an endless loop that cancels a random request and
// creates a new one every 200 ms
func (app *App) SimulateChanges() {
	for {
		time.Sleep(200 * time.Millisecond)
		app.CancelRandomRequest()
		app.CreateRandomRequest()
	}
}
