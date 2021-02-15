package cli

import (
	"fmt"
	"strconv"

	"github.com/jorgeAM/brew/internal/domain"

	"github.com/spf13/cobra"
)

type cobraFn func(cmd *cobra.Command, args []string) error

func InitBeerCmd(repository domain.Repository) *cobra.Command {
	beerCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Aliases: []string{
			"b",
		},
		RunE: runBeersFn(repository),
	}

	beerCmd.Flags().StringP("id", "i", "", "id of beer")

	return beerCmd
}

func runBeersFn(repository domain.Repository) cobraFn {
	return func(cmd *cobra.Command, args []string) error {
		beers, err := repository.GetBeers()

		if err != nil {
			return err
		}

		id, _ := cmd.Flags().GetString("id")

		if id != "" {
			i, _ := strconv.Atoi(id)
			for _, beer := range beers {
				if beer.ID == i {
					fmt.Println(beer)
					return nil
				}
			}
		} else {
			fmt.Println(beers)
		}

		return nil
	}
}
