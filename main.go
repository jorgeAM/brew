package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var favoriteBeers = map[string]string{
	"b1": "7 vidas chocolate",
	"b2": "The red moon",
	"b3": "barbarian",
}

func main() {
	beerCmd := flag.NewFlagSet("beers", flag.ExitOnError)
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("You must specified a commad beers")
		os.Exit(2)
	}

	switch flag.Arg(0) {
	case "beers":
		ID := beerCmd.String("id", "", "find by ID")
		beerCmd.Parse(os.Args[2:])

		if *ID != "" {
			fmt.Println("ID: ", *ID)
			fmt.Println(favoriteBeers[*ID])
		} else {
			fmt.Println(favoriteBeers)
		}
	}
}
