package config

import "os"

type NewRelicConfig struct {
	LicenseKey 		string
	AppName    		string 
}

func GetNewRelicConfig() *NewRelicConfig{
	return &NewRelicConfig{
		LicenseKey: os.Getenv("NEWRELIC_LICENSE_KEY"),
		AppName: "fortune",
	}
}