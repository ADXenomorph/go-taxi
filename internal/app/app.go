package app

import (
	"errors"
	"time"

	"github.com/ADXenomorph/go-taxi/internal/taxi_request"
)

type App struct {
	storage *taxi_request.RequestStorage
}

func NewApp(storage *taxi_request.RequestStorage) *App {
	return &App{storage: storage}
}

func (app *App) CreateRequest(id string) *taxi_request.Request {
	tr := taxi_request.NewRequest(id)
	app.storage.Save(tr)

	return tr
}

func (app *App) CancelRequest(id string) error {
	tr, ok := app.storage.Get(id)

	if !ok {
		return errors.New("Failed to cancel")
	}

	tr.Status = taxi_request.Cancelled
	app.storage.Save(tr)

	return nil
}

func (app *App) GetRandomRequest() *taxi_request.Request {
	return app.storage.GetRandomAndCount()
}

func (app *App) GetRequestStatistics() []string {
	return app.storage.GetCounters()
}

func (app *App) CreateRandomRequest() *taxi_request.Request {
	return app.CreateRequest(taxi_request.GenerateRequestId())
}

func (app *App) CreateInitialRequests() {
	for i := 1; i <= 50; i++ {
		app.CreateRandomRequest()
	}
}

func (app *App) CancelRandomRequest() {
	req := app.storage.GetRandom()
	app.CancelRequest(req.ID)
}

func (app *App) SimulateChanges() {
	for {
		time.Sleep(200 * time.Millisecond)
		app.CancelRandomRequest()
		app.CreateRandomRequest()
	}
}
