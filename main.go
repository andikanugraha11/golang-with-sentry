package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/andikanugraha11/golang-with-sentry/config"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config.InitSentry()
	defer sentry.Flush(2 * time.Second)

	// Create an instance of sentryhttp
	sentryMiddleware := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(sentryMiddleware.Handle)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
		sentry.CaptureException(errors.New("welcome"))
	})
	r.Get("/warning", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("warning"))
		sentry.CaptureMessage("test")
	})
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("server panic")
	})
	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		hub := sentry.GetHubFromContext(r.Context())
		hub.CaptureException(errors.New("test error"))
	})

	log.Println("Server Running On port 3000")
	http.ListenAndServe(":3000", r)
}
