package repo

type Repository interface {
	All() ([]OrderData, error)
	GetById(id string) (*OrderData, error)
	Insert(OrderData) error
}
