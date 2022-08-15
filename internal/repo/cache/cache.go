package cache

import (
	"errors"
	"wb_l0/internal/repo"
)

type Repository struct {
	Cache map[string]repo.OrderData
}

func NewCacheRepository() *Repository {
	return &Repository{
		Cache: map[string]repo.OrderData{},
	}
}

//func (r *Repository) Create(orderData repo.OrderData) (*repo.OrderData, error) {
//	id := orderData.OrderUid
//	(*r).Cache[id] = orderData
//	return &orderData, nil
//}

func (r *Repository) All() ([]repo.OrderData, error) {
	var allRecords []repo.OrderData
	for _, od := range (*r).Cache {
		allRecords = append(allRecords, od)
	}
	return allRecords, nil
}

func (r *Repository) GetById(id string) (*repo.OrderData, error) {
	if record, ok := (*r).Cache[id]; ok {
		return &record, nil
	} else {
		return nil, errors.New("getting record by id: this order_id does not exist in cache")
	}
}

func (r *Repository) Insert(record repo.OrderData) error {
	(*r).Cache[record.OrderUid] = record
	return nil
}
