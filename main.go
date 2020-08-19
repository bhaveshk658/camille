package main

import (
	"flag"
	"fmt"

	"github.com/gocolly/colly"
)

type Text struct {
	Info string
}

const baseURL = "https://blitz.gg/lol/champions/Camille/counters/"

func getWinRate(fetchURL string) Text {

	texts := make([]Text, 0)

	c := colly.NewCollector()

	/*
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting: ", r.URL)
		})
	*/

	c.OnHTML(`.ChampionMatchupStatsHeader__Caption-sc-16vko7r-0`, func(e *colly.HTMLElement) {
		text := Text{Info: e.Text}
		texts = append(texts, text)
	})

	c.Visit(fetchURL)
	if len(texts) > 0 {
		return texts[0]
	}
	return Text{Info: "No matchup information available"}
}

func main() {
	matchup := flag.String("vs", "None", "Display matchup information")
	flag.Parse()
	if *matchup == "None" {
		return
	}
	fetchURL := baseURL + *matchup
	winrate := getWinRate(fetchURL)
	fmt.Println(winrate.Info)
}
