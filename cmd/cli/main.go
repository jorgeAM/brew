package main

import (
	"fmt"
	"os"

	"github.com/jorgeAM/brew/internal/cli"
	"github.com/jorgeAM/brew/internal/storage/csv"

	"github.com/spf13/cobra"
)

func main() {
	repository := csv.NewRepository()
	// repository := otario.NewRepository()
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeerCmd(repository))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
