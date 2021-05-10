package addCart

import (
	"encoding/json"
	"log"
	"net/http"
)

type Store interface {
	Get() ([]cart, error)
	Add(c cart) error
}

type cart struct {
	User     int64
	Product  string
	Quantity int
}

func AddToCart(w http.ResponseWriter, r *http.Request, s Store) {
	decoder := json.NewDecoder(r.Body)
	var t cart
	err := decoder.Decode(&t)
	if err != nil {
		log.Fatalf("Unable to add to cart.", err)
	}
	c := cart{t.User, t.Product, t.Quantity}
	s.Add(c)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Added to cart")
}

// Get all carts.
func GetData(w http.ResponseWriter, r *http.Request, s Store) {
	carts, err := s.Get()
	if err != nil {
		log.Fatalf("Unable to get all pending carts", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(carts)
}
