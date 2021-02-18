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
	/* Profiling CPU */
	// f2, _ := os.Create("beers.cpu.prof")
	// defer f2.Close()
	// pprof.StartCPUProfile(f2)
	// defer pprof.StopCPUProfile()

	var repository domain.Repository

	csvData := flag.Bool("csv", false, "use csv repository")
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

	/* Profiling memoria */
	// f2, _ := os.Create("beers.mem.prof")
	// defer f2.Close()
	// pprof.WriteHeapProfile(f2)

}
