package config

import (
	"log"

	"github.com/getsentry/sentry-go"
)

func InitSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://examplePublicKey@o0.ingest.sentry.io/0",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
