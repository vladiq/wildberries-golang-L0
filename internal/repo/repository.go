package repo

type Repository interface {
	Create(orderData OrderData) (*OrderData, error)
	All() ([]OrderData, error)
	GetById(id string) (*OrderData, error)
	Insert(OrderData) error
}
