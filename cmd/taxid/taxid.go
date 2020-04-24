// A simple taxi-like app, that serves as a high-load app demonstration using golang.
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"github.com/ADXenomorph/go-taxi/internal/taxi"
	"github.com/ADXenomorph/go-taxi/internal/taxi_request"
)

func main() {
	app := taxi.NewApp(taxi_request.NewStorage())

	// Create initial 50 taxi requests
	app.CreateInitialRequests()

	// Constantly cancel and add new requests to simulate incoming changes
	go app.SimulateChanges()

	router := CreateRouter(app)

	// Start the web server
	log.Print("Starting HTTP server on :8080")
	go func() {
		if err := fasthttp.ListenAndServe(":8080", router.Handler); err != nil {
			log.Fatalf("error in ListenAndServe: %s", err)
		}
	}()

	// Wait forever
	select {}
}

// CreateRouter sets up API mapping handlers and returns the router.
func CreateRouter(app *taxi.App) *fasthttprouter.Router {
	r := fasthttprouter.New()

	r.GET("/request", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprint(ctx, app.GetRandomRequest().ID)
	})
	r.GET("/admin/requests", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprint(ctx, strings.Join(app.GetRequestStatistics(), "\n"))
	})

	return r
}
