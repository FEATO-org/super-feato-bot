package config

import (
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var SCOOPS []string

func init() {
	SCOOPS = append(SCOOPS, "https://www.googleapis.com/auth/drive", "https://www.googleapis.com/auth/spreadsheets")
}

func NewOauth2() *oauth2.Config {
	// FIXME: load env
	credentials, err := os.ReadFile("./credentials.json")
	if err != nil {
		log.Printf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(credentials, SCOOPS...)
	if err != nil {
		log.Printf("Unable to parse client secret file to config: %v", err)
	}

	return config
}
