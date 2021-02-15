package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type cobraFn func(cmd *cobra.Command, args []string) error

func InitBeerCmd() *cobra.Command {
	beerCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Aliases: []string{
			"b",
		},
		RunE: runBeersFn(),
	}

	beerCmd.Flags().StringP("id", "i", "", "id of beer")

	return beerCmd
}

func runBeersFn() cobraFn {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Println("beer was call")

		file, err := os.Open("./data/beers.csv")

		if err != nil {
			return err
		}

		reader := bufio.NewReader(file)

		var beers = make(map[int]string)
		for line := readLine(reader); line != nil; line = readLine(reader) {
			values := strings.Split(string(line), ",")

			productID, _ := strconv.Atoi(values[0])
			beers[productID] = values[1]
		}

		id, _ := cmd.Flags().GetString("id")

		if id != "" {
			i, _ := strconv.Atoi(id)
			fmt.Println(beers[i])
		} else {
			fmt.Println(beers)
		}

		return nil
	}
}

func readLine(reader *bufio.Reader) []byte {
	line, _, _ := reader.ReadLine()
	return line
}
