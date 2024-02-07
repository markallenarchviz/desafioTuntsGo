package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	// create api context
	ctx := context.Background()

	jsonData, err := ioutil.ReadFile("google.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Encode JSON data to Base64
	base64Data := base64.StdEncoding.EncodeToString(jsonData)

	// get bytes from base64 encoded google service accounts key
	credBytes, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		log.Fatal(err)
		return
	}

	// authenticate and get configuration
	config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatal(err)
		return
	}

	// create client with config and context
	client := config.Client(ctx)

	// create new service using client
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatal(err)
		return
	}

	spreadsheetId := "1jJCYe__gprZznPm-iecd4x5p0W9oXxXyhlHOqhF7IAY"
	var rang = "A4:H27"

	// Convert sheet ID to sheet name.
	response1, err := srv.Spreadsheets.Values.Get(spreadsheetId, rang).Do()
	if err != nil || response1.HTTPStatusCode != 200 {
		log.Fatal(err)
		return
	}

	fmt.Println(response1)

}
