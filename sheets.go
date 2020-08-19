package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

const spreadsheetID = "1C5scmVWQ42MePRg6AZr4_DTy5Xc5WvdEhiiM0Eszjt4"

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	data, err := ioutil.ReadFile("secret.json")
	checkError(err)
	conf, err := google.JWTConfigFromJSON(data, sheets.SpreadsheetsScope)
	checkError(err)

	client := conf.Client(context.TODO())
	srv, err := sheets.New(client)
	checkError(err)

	readRange := "A1"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	checkError(err)

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			fmt.Printf("%s\n", row[0])
		}
	}

}
