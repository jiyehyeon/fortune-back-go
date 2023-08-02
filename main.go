package main

import (
	"fmt"
	"fortune-back-go/pkg/config"
	"fortune-back-go/pkg/middleware"
	"fortune-back-go/pkg/router"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading test.env file:", err)
		return
	}

	newRelic := config.GetNewRelicConfig()

	app, err := newrelic.NewApplication(
    newrelic.ConfigLicense(newRelic.LicenseKey),
    newrelic.ConfigAppName(newRelic.AppName),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	if err != nil{
		log.Fatalf("newRelic err: %v", err)
	}

	mux := http.NewServeMux()
	fortuneRouter := router.NewFortuneRouter().Initialize()

	mux.Handle(newrelic.WrapHandle(app, "/fortune", middleware.JSONHeader(fortuneRouter)))

	fmt.Println("server starting..")
	http.ListenAndServe(":8080", mux)
}
