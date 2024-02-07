package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

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
	allSheets, err := srv.Spreadsheets.Values.Get(spreadsheetId, rang).Do()
	if err != nil {
		log.Fatal(err)
		return
	}

	size := len(allSheets.Values)

	var result []any

	for i := 0; i < size; i++ {
		absence, _ := allSheets.Values[i][2].(string)
		p1, _ := allSheets.Values[i][3].(string)
		p2, _ := allSheets.Values[i][4].(string)
		p3, _ := allSheets.Values[i][5].(string)
		res := studentStatus(absence, p1, p2, p3)
		arr := []string{res, finalExame(res, p1, p2, p3)}
		result = append(result, arr)
	}

	fmt.Println(result)
	res := studentStatus("50", "66", "88", "80")
	fmt.Println(res)
}

func studentStatus(absence, p1, p2, p3 string) string {
	absenceInt, _ := strconv.ParseFloat(absence, 64)
	p1Int, _ := strconv.Atoi(p1)
	p2Int, _ := strconv.Atoi(p2)
	p3Int, _ := strconv.Atoi(p3)
	absences := int((absenceInt / 60) * 100)
	average := ((p1Int / 10) + (p2Int / 10) + (p3Int / 10)) / 3
	if absences >= 25 {
		return "Reprovado por Falta"
	}

	if average >= 7 {
		return "Aprovado"
	} else if average < 7 && average >= 5 {
		return "Exame Final"
	} else {
		return "Reprovado por Nota"
	}
}

func finalExame(res, p1, p2, p3 string) string {
	p1Int, _ := strconv.ParseFloat(p1, 64)
	p2Int, _ := strconv.ParseFloat(p2, 64)
	p3Int, _ := strconv.ParseFloat(p3, 64)

	average := int(((p1Int / 10) + (p2Int / 10) + (p3Int / 10)) / 3)

	if res != "Exame Final" {
		return "0"
	}

	return strconv.Itoa(((average + 5) / 2))
}
