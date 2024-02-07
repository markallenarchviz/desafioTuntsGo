package main

import (
	"context"
	"log"

	"github.com/markallenarchviz/desafioTuntsGo/config"
	"github.com/markallenarchviz/desafioTuntsGo/pkg"
	"google.golang.org/api/sheets/v4"
)

func main() {

	ctx := context.Background()

	srv := config.Config()

	spreadsheetId := "1jJCYe__gprZznPm-iecd4x5p0W9oXxXyhlHOqhF7IAY"
	var rang = "A4:H27"

	// Convert sheet ID to sheet name.
	allSheets, err := srv.Spreadsheets.Values.Get(spreadsheetId, rang).Do()
	if err != nil {
		log.Fatal(err)
		return
	}

	size := len(allSheets.Values)

	var result [][]interface{}

	//Loop over spreadsheet to add new values
	for i := 0; i < size; i++ {
		absence, _ := allSheets.Values[i][2].(string)
		p1, _ := allSheets.Values[i][3].(string)
		p2, _ := allSheets.Values[i][4].(string)
		p3, _ := allSheets.Values[i][5].(string)
		res := pkg.StudentStatus(absence, p1, p2, p3)
		arr := []interface{}{res, pkg.FinalExames(res, p1, p2, p3)}
		result = append(result, arr)
	}

	range2 := "G4:H27"

	//Creating rows data type
	row := &sheets.ValueRange{
		Values: result,
	}

	//Update the spreadsheet with new data
	_, err = srv.Spreadsheets.Values.Update(spreadsheetId, range2, row).ValueInputOption("USER_ENTERED").Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
}
