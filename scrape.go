package main

import "github.com/gocolly/colly"

const baseURL = "https://blitz.gg/lol/champions/Camille/counters/"

type Text struct {
	Info string
}

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
