package main_test

import (
	"testing"
	"time"

	main "github.com/ADXenomorph/go-taxi/cmd/taxid"
	taxi "github.com/ADXenomorph/go-taxi/internal/app"
	"github.com/ADXenomorph/go-taxi/internal/taxi_request"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func TestAppServer(t *testing.T) {
	app := taxi.NewApp(taxi_request.NewStorage())

	router := main.CreateRouter(app)

	ln := fasthttputil.NewInmemoryListener()

	serverCh := make(chan struct{})
	go func() {
		if err := fasthttp.Serve(ln, router.Handler); err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		close(serverCh)
	}()

	clientCh := make(chan struct{})
	go func() {
		c, err := ln.Dial()
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if _, err = c.Write([]byte("GET /request HTTP/1.1\r\nHost: aa\r\n\r\n")); err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if err = c.Close(); err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		close(clientCh)
	}()

	select {
	case <-clientCh:
	case <-time.After(time.Second):
		t.Fatal("timeout")
	}

	if err := ln.Close(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	select {
	case <-serverCh:
	case <-time.After(time.Second):
		t.Fatal("timeout")
	}
}
