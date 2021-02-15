package domain

type Beer struct {
	ID       int
	Name     string
	Price    float64
	Category string
	Type     *BeerType
	Brewer   string
	Country  string
}

func NewBeer(id int, name, category, brewer, country string, beerType *BeerType, price float64) Beer {
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
