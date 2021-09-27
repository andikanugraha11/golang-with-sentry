package config

import (
	"log"

	"github.com/getsentry/sentry-go"
)

func InitSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "<URL_SENTRY>",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	sentry.CaptureMessage("It works!")
}
