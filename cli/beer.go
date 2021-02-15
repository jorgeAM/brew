package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type cobraFn func(cmd *cobra.Command, args []string)

var favoriteBeers = map[string]string{
	"b1": "7 vidas chocolate",
	"b2": "The red moon",
	"b3": "barbarian",
}

func InitBeerCmd() *cobra.Command {
	beerCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Aliases: []string{
			"b",
		},
		Run: runBeersFn(),
	}

	beerCmd.Flags().StringP("id", "i", "", "id of beer")

	return beerCmd
}

func runBeersFn() cobraFn {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("beer was call")

		id, _ := cmd.Flags().GetString("id")

		if id != "" {
			fmt.Println(favoriteBeers[id])
		} else {
			fmt.Println(favoriteBeers)
		}
	}
}
