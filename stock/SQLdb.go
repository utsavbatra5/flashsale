package stock

import (
	"flashsale/util"
	"sync"
)

type SQLdb struct {
	mu         sync.Mutex
	totalStock map[string]int
}

func NewSQLdb(config util.DBconfig) Store {
	return nil
}

// AddQuantity adds quantity/ stock of product.
func (m *SQLdb) AddQuantity(qtyData qtyData) {
	// Code here
}

// Get Product data
func (m *SQLdb) Get(product string) (qtyData, error) {
	// Code here
	return qtyData{}, nil
}

// UpdateQuantity updates quantity
func (m *SQLdb) UpdateQuantity(q qtyData) error {
	// Code here
	return nil
}
