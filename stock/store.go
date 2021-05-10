package stock

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Store interface {
	Get(string) (qtyData, error)
	UpdateQuantity(data qtyData) error
}

type qtyData struct {
	Product  string
	Quantity int
}

func GetStock(w http.ResponseWriter, r *http.Request, s Store) {
	vars := mux.Vars(r)
	productID := vars["prodID"]
	qty, err := s.Get(productID)
	if err != nil {
		log.Fatalf("Unable to fetch quantity.", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(qty)
}

func UpdateStock(w http.ResponseWriter, r *http.Request, s Store) {
	decoder := json.NewDecoder(r.Body)
	var q qtyData
	err := decoder.Decode(&q)
	if err != nil {
		log.Fatalf("Unable to update stock", err)
	}
	err = s.UpdateQuantity(q)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Quantity Updated")
}
