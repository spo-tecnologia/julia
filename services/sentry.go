package services

import (
	"fmt"
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
)

func InitSentry() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:           sentryDSN,
			EnableTracing: false,
		}); err != nil {
			fmt.Printf("Sentry initialization failed: %v", err)
		}
		fmt.Println("Sentry initialized")
	}

}
