package stock

import (
	"sync"
)

type InMem struct {
	Product string
	Stock   int
}
type InMemdb struct {
	mu         sync.Mutex
	totalStock map[string]int
}

func NewInMemdb() Store {
	m := make(map[string]int)
	m["X"] = 10
	return &InMemdb{sync.Mutex{}, m}
}

// AddQuantity adds quantity/ stock of product.
func (m *InMemdb) AddQuantity(qtyData qtyData) {
	m.mu.Lock()
	m.totalStock[qtyData.Product] = qtyData.Quantity
	m.mu.Unlock()
}

// Get Product data
func (m *InMemdb) Get(product string) (qtyData, error) {
	return qtyData{product, m.totalStock[product]}, nil
}

// UpdateQuantity updates quantity
func (m *InMemdb) UpdateQuantity(q qtyData) error {
	m.mu.Lock()
	m.totalStock[q.Product] = m.totalStock[q.Product] + q.Quantity
	m.mu.Unlock()
	return nil
}
