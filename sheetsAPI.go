package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

const spreadsheetID = "1C5scmVWQ42MePRg6AZr4_DTy5Xc5WvdEhiiM0Eszjt4"

type Matchup struct {
	Name         string
	Difficulty   string
	AbilityStart string
	Tips         string
	Runes        string
	Changes      string
	ItemStart    string
	ItemRush     string
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func getMatchup(name *string) Matchup {

	data, err := ioutil.ReadFile("secret.json")
	checkError(err)
	conf, err := google.JWTConfigFromJSON(data, sheets.SpreadsheetsScope)
	checkError(err)

	client := conf.Client(context.TODO())
	srv, err := sheets.New(client)
	checkError(err)

	readRange := "A2:H57"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	checkError(err)

	matchup := Matchup{Name: "None"}

	if len(resp.Values) == 0 {
		fmt.Println("Matchup info not found.")
	} else {
		for _, row := range resp.Values {
			if row[0] == *name {

				matchup = Matchup{
					Name:         row[0].(string),
					Difficulty:   row[1].(string),
					AbilityStart: row[2].(string),
					Tips:         row[3].(string),
					Runes:        row[4].(string),
					Changes:      row[5].(string),
					ItemStart:    row[6].(string),
					ItemRush:     row[7].(string),
				}

				return matchup
			}
		}
	}

	return matchup

}
