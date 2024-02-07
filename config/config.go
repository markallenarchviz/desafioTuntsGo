package config

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func Config() *sheets.Service {
	// create api context
	ctx := context.Background()

	jsonData, err := os.ReadFile("google.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Encode JSON data to Base64
	base64Data := base64.StdEncoding.EncodeToString(jsonData)

	// get bytes from base64 encoded google service accounts key
	credBytes, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// authenticate and get configuration
	config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// create client with config and context
	client := config.Client(ctx)

	// create new service using client
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return srv

}
