package domain

import "encoding/json"

type BeerType int

const (
	Unknown BeerType = iota
	Lager
	Malt
	Ale
	FlavouredMalt
	Stout
	Porter
	NonAlcoholic
)

func NewBeerType(t string) *BeerType {
	beerType := toID[t]
	return &beerType
}

func (t BeerType) String() string {
	return toString[t]
}

func (t *BeerType) UnmarshalJSON(bytes []byte) error {
	var j string

	err := json.Unmarshal(bytes, &j)

	if err != nil {
		return err
	}

	*t = toID[j]

	return nil
}

var toString = map[BeerType]string{
	Lager:         "Lager",
	Malt:          "Malt",
	Ale:           "Ale",
	FlavouredMalt: "Flavoured Malt",
	Stout:         "Stout",
	Porter:        "Stout",
	NonAlcoholic:  "Non-Alcoholic",
	Unknown:       "unknown",
}

var toID = map[string]BeerType{
	"Lager":          Lager,
	"Malt":           Malt,
	"Ale":            Ale,
	"Flavoured Malt": FlavouredMalt,
	"Stout":          Stout,
	"Porter":         Porter,
	"Non-Alcoholic":  NonAlcoholic,
	"unknown":        Unknown,
}
