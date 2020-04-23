package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	taxi "github.com/ADXenomorph/go-taxi/internal/app"
	"github.com/ADXenomorph/go-taxi/internal/taxi_request"
)

func main() {
	app := taxi.NewApp(taxi_request.NewStorage())
	app.CreateInitialRequests()

	go app.SimulateChanges()

	router := CreateRouter(app)

	log.Print("Starting HTTP server on :8080")
	go func() {
		if err := fasthttp.ListenAndServe(":8080", router.Handler); err != nil {
			log.Fatalf("error in ListenAndServe: %s", err)
		}
	}()

	// Wait forever
	select {}
}

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
