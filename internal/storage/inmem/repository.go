package inmem

import "github.com/jorgeAM/brew/internal/domain"

type repository struct{}

func NewRepository() domain.Repository {
	return &repository{}
}

func (r *repository) GetBeers() ([]domain.Beer, error) {
	return []domain.Beer{
		domain.NewBeer(
			127,
			"Mad Jack Mixer",
			"Domestic Specialty",
			"Molson",
			"Canada",
			"23.95",
			domain.NewBeerType("Lager"),
		),
		domain.NewBeer(
			8520130,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Grolsch Export B.V.",
			"Canada",
			"49.50",
			domain.NewBeerType("Non-Alcoholic Beer"),
		),
	}, nil
}
