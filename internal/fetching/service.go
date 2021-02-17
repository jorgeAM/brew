package fetching

import (
	"github.com/jorgeAM/brew/internal/domain"
	"github.com/pkg/errors"
)

type Service interface {
	FetchBeers() ([]domain.Beer, error)
	FetchByID(id int) (domain.Beer, error)
}

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) Service {
	return &service{r}
}

func (s *service) FetchBeers() ([]domain.Beer, error) {
	return s.repository.GetBeers()
}

func (s *service) FetchByID(id int) (domain.Beer, error) {
	beers, err := s.FetchBeers()
	if err != nil {
		return domain.Beer{}, err
	}

	for _, beer := range beers {
		if beer.ID == id {
			return beer, nil
		}
	}

	return domain.Beer{}, errors.Errorf("Beer %d not found", id)
}
