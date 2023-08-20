package main

import (
	"fmt"
	"fortune-back-go/pkg/config"
	"fortune-back-go/pkg/middleware"
	"fortune-back-go/pkg/router"
	"log"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
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
	mux.Handle(newrelic.WrapHandle(app, "/fortune/", middleware.JSONHeader(http.StripPrefix("/fortune", fortuneRouter))))

	fmt.Println("server starting..")
	http.ListenAndServe(":8080", mux)
}
