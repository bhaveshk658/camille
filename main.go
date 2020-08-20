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

	matchup := GetMatchup(matchupName)

	fmt.Println(winrate.Info)
	fmt.Printf("Difficulty: %s\n", matchup.Difficulty)
	fmt.Printf("Ability Start: %s\n", matchup.AbilityStart)
	fmt.Printf("Tips:\n %s\n", matchup.Tips)
	fmt.Printf("Runes: %s\n", matchup.Runes)
	fmt.Printf("Changes: %s\n", matchup.Changes)
	fmt.Printf("Starting Item: %s\n", matchup.ItemStart)
	fmt.Printf("Item Rush: %s\n", matchup.ItemRush)
}
