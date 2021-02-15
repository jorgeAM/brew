package domain

type Beer struct {
	ID       int       `json:"beer_id"`
	Name     string    `json:"name"`
	Price    string    `json:"price"`
	Category string    `json:"category"`
	Type     *BeerType `json:"-"`
	Brewer   string    `json:"brewer"`
	Country  string    `json:"country"`
}

func NewBeer(id int, name, category, brewer, country, price string, beerType *BeerType) Beer {
	beer := Beer{
		ID:       id,
		Name:     name,
		Category: category,
		Type:     beerType,
		Brewer:   brewer,
		Country:  country,
		Price:    price,
	}

	return beer
}
