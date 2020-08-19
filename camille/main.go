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
	matchupName := flag.String("vs", "None", "Display matchup information")
	flag.Parse()
	if *matchupName == "None" {
		return
	}
	fetchURL := baseURL + *matchupName
	winrate := getWinRate(fetchURL)

	matchup := camille.getMatchup(matchupName)

	fmt.Printf("Winrate: %s\n", winrate.Info)
	fmt.Printf("Difficulty: %s\n", matchup.Difficulty)
}
