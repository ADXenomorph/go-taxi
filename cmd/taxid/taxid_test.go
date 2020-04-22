package main_test

import (
	"testing"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"

	main "github.com/ADXenomorph/go-taxi/cmd/taxid"
	taxi "github.com/ADXenomorph/go-taxi/internal/app"
	"github.com/ADXenomorph/go-taxi/internal/taxi_request"
)

var router *fasthttprouter.Router
var ln *fasthttputil.InmemoryListener

func TestMain(m *testing.M) {
	app := taxi.NewApp(taxi_request.NewStorage())
	app.CreateInitialRequests()

	go app.SimulateChanges()

	router = main.CreateRouter(app)

	ln := fasthttputil.NewInmemoryListener()
	fasthttp.Serve(ln, router.Handler)

	m.Run()
}

func BenchmarkMainParallel(b *testing.B) {
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			makeRequest(b)
		}
	})
}

func makeRequest(b *testing.B) {
	c, err := ln.Dial()
	if err != nil {
		b.Errorf("unexpected error: %s", err)
	}
	if _, err = c.Write([]byte("GET /request HTTP/1.1\r\nHost: localhost:8080\r\n\r\n")); err != nil {
		b.Errorf("unexpected error: %s", err)
	}
}
