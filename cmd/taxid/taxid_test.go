package main_test

import (
	"log"
	"runtime"
	"testing"
	"time"

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

	ln = fasthttputil.NewInmemoryListener()

	serverCh := make(chan struct{})
	go func() {
		if err := fasthttp.Serve(ln, router.Handler); err != nil {
			fail("unexpected error: %s", err)
		}
		close(serverCh)
	}()

	m.Run()

	select {
	case <-serverCh:
	case <-time.After(time.Second):
		fail("server timeout")
	}

	// if err := ln.Close(); err != nil {
	// 	fail("unexpected error: %s", err)
	// }
}

func fail(format string, v ...interface{}) {
	log.Fatalf(format, v...)
	runtime.Goexit()
}

func BenchmarkMainParallel(b *testing.B) {
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			makeRequest(ln)
		}
	})
}

func TestStuff(t *testing.T) {
	clientCh := make(chan struct{})
	go func() {
		makeRequest(ln)
		close(clientCh)
	}()

	select {
	case <-clientCh:
	case <-time.After(time.Second):
		t.Fatal("timeout")
	}
}

func makeRequest(ln *fasthttputil.InmemoryListener) {
	c, err := ln.Dial()
	if err != nil {
		fail("unexpected error: %s", err)
	}
	if _, err = c.Write([]byte("GET /request HTTP/1.1\r\nHost: localhost:8080\r\n\r\n")); err != nil {
		fail("unexpected error: %s", err)
	}
	if err = c.Close(); err != nil {
		fail("unexpected error: %s", err)
	}
	log.Println("make request")
}
