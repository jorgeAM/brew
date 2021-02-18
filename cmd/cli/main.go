package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jorgeAM/brew/internal/cli"
	"github.com/jorgeAM/brew/internal/domain"
	"github.com/jorgeAM/brew/internal/fetching"
	"github.com/jorgeAM/brew/internal/storage/csv"
	"github.com/jorgeAM/brew/internal/storage/otario"

	"github.com/spf13/cobra"
)

func main() {
	var repository domain.Repository

	csvData := flag.Bool("csv", true, "use csv repository")
	flag.Parse()

	repository = otario.NewRepository()

	if *csvData == true {
		repository = csv.NewRepository()
	}

	fmt.Println(*csvData)

	service := fetching.NewService(repository)

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeerCmd(service))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
