package main

import (
	"flag"
	"fmt"
)

type Text struct {
	Info string
}

func main() {
	matchupName := flag.String("vs", "None", "Display matchup information")
	flag.Parse()
	if *matchupName == "None" {
		return
	}
	fetchURL := baseURL + *matchupName
	winrate := getWinRate(fetchURL)

	matchup := getMatchup(matchupName)

	fmt.Println(winrate.Info)
	fmt.Printf("Difficulty: %s\n", matchup.Difficulty)
	fmt.Printf("Ability Start: %s\n", matchup.AbilityStart)
	fmt.Printf("Tips:\n %s\n", matchup.Tips)
	fmt.Printf("Runes: %s\n", matchup.Runes)
	fmt.Printf("Changes: %s\n", matchup.Changes)
	fmt.Printf("Starting Item: %s\n", matchup.ItemStart)
	fmt.Printf("Item Rush: %s\n", matchup.ItemRush)
}
