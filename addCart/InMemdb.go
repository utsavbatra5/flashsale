package addCart

import (
	"sync"
)

//var (
//	ourCart = SQLdb{productStock: make(map[string]int)}
//)

//type userQty struct{
//	user int64
//	product string
//	quantity int
//}
//type allUserQty []userQty

type InMem struct {
	mu           sync.Mutex
	productStock []cart
}

func NewInMemdb() Store {
	return &InMem{sync.Mutex{}, []cart{}}
}

func GetInMem() Store {
	return &InMem{sync.Mutex{}, []cart{}}
}

// AddQuantity adds quantity/ stock of product.
func (m *InMem) AddQuantity(user int64, product string, quantity int) {
	uc := cart{user, product, quantity}
	m.mu.Lock()
	m.productStock = append(m.productStock, uc)
	m.mu.Unlock()
}

// Get all carts.
func (m *InMem) Get() ([]cart, error) {
	return m.productStock, nil
}

// Add inserts new cart to memory.
func (m *InMem) Add(c cart) error {
	m.AddQuantity(c.User, c.Product, c.Quantity)
	return nil
}
