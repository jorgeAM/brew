package otario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jorgeAM/brew/internal/domain"
)

const (
	productsEndpoint = "/products"
	ontarioURL       = "http://ontariobeerapi.ca"
)

type repository struct {
	url string
}

func NewRepository() domain.Repository {
	url := fmt.Sprintf("%v%v", ontarioURL, productsEndpoint)

	return &repository{url}
}

func (r *repository) GetBeers() ([]domain.Beer, error) {
	res, err := http.Get(r.url)

	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var beers []domain.Beer

	err = json.Unmarshal(bytes, &beers)

	if err != nil {
		return nil, err
	}

	return beers, nil
}
