package addCart

import (
	"flashsale/util"
	"sync"
)

type SQLdb struct {
	mu           sync.Mutex
	productStock []cart
}

func NewSQLdb(config util.DBconfig) Store {
	return nil
}

func GetSQLdb() Store {
	return nil
}

// AddQuantity adds quantity/ stock of product.
func (m *SQLdb) AddQuantity(user int64, product string, quantity int) {
	//Code here.
}

// Get all carts.
func (m *SQLdb) Get() ([]cart, error) {
	return nil, nil
}

// Add inserts new cart to memory.
func (m *SQLdb) Add(c cart) error {
	return nil
}
