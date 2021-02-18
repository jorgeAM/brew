package cli

import (
	"fmt"
	"strconv"

	"github.com/jorgeAM/brew/internal/fetching"

	"github.com/spf13/cobra"
)

type cobraFn func(cmd *cobra.Command, args []string) error

func InitBeerCmd(service fetching.Service) *cobra.Command {
	beerCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Aliases: []string{
			"b",
		},
		RunE: runBeersFn(service),
	}

	beerCmd.Flags().StringP("id", "i", "", "id of beer")

	return beerCmd
}

func runBeersFn(service fetching.Service) cobraFn {
	return func(cmd *cobra.Command, args []string) error {
		id, _ := cmd.Flags().GetString("id")

		if id == "" {
			beers, err := service.FetchBeers()

			if err != nil {
				return err
			}

			fmt.Println(beers)

			return nil
		}

		i, err := strconv.Atoi(id)

		if err != nil {
			return err
		}

		beer, err := service.FetchByID(i)

		if err != nil {
			return err
		}

		fmt.Println(beer)

		return nil
	}
}
