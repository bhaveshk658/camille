package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

var sentence = []string{"Camille vs"}

func info() {
	app.Name = "Camille Matchups Guide"
	app.Usage = "A CLI for Camille mains to learn matchups"
	author := cli.Author{Name: "bhaveshk658", Email: "bhaveshk658@berkeley.edu"}
	app.Authors = append(app.Authors, &author)
	app.Version = "1.0.0"

}

func commands() {
	app.Commands = []*cli.Command{
		{
			Name:    "renekton",
			Aliases: []string{"croc"},
			Usage:   "Print Renekton matchup info",
			Action: func(c *cli.Context) error {
				opponent := "Renekton"
				matchup := append(sentence, opponent)
				m := strings.Join(matchup, " ")
				fmt.Println(m)
				return nil
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
