package domain

type Repository interface {
	GetBeers() ([]Beer, error)
}
