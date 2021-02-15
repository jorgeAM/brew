package csv

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/jorgeAM/brew/internal/domain"
)

type repository struct{}

func NewRepository() domain.Repository {
	return &repository{}
}

func (r *repository) GetBeers() ([]domain.Beer, error) {
	file, err := os.Open("./data/beers.csv")

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	var beers []domain.Beer

	for line := readLine(reader); line != nil; line = readLine(reader) {
		values := strings.Split(string(line), ",")

		id, _ := strconv.Atoi(values[0])

		beer := domain.NewBeer(
			id,
			values[1],
			values[2],
			values[5],
			values[6],
			values[3],
			domain.NewBeerType(values[4]),
		)

		beers = append(beers, beer)
	}

	return beers, nil
}

func readLine(reader *bufio.Reader) []byte {
	line, _, _ := reader.ReadLine()
	return line
}
