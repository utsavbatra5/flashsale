package shop

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ProductStock struct {
	Product  string
	Quantity int
}

var (
	getStockURL    = "http://localhost:8001/api/"
	updateStockURL = "http://localhost:8001/api/updatestock"
	basic          = "Basic dXNlcjpwYXNz"
)

// checkStock checks current stock of the product
func checkStock(product string) int {
	url := getStockURL + product
	req, err := http.NewRequest("GET", url, nil)
	// add authorization header to the req
	req.Header.Add("Authorization", basic)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error on response", err)
	}
	defer resp.Body.Close()

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var productStock ProductStock
	err1 := json.Unmarshal(body, &productStock)
	//Convert the body to type string
	if err1 != nil {
		log.Fatalln(err1)
	}

	return productStock.Quantity

}

// updateStock updates stock of the product
func updateStock(ps ProductStock) error {
	values := ps

	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("POST", updateStockURL, bytes.NewBuffer(jsonValue))
	// add authorization header to the req
	req.Header.Add("Authorization", basic)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error on response", err)
	}
	defer resp.Body.Close()

	return nil
}

// triggerPayments triggers payment gateway for the product. Dummy implementation of payment gateway provided.
func triggerPayment(ch chan TransactionData, sem semaphore) {
	defer sem.release(1)
	data, err := getVal()
	if err != nil {

	}
	go executePayment(data, ch)
}

// BuyProducts initiates checkout and payment of the product
func BuyProducts(w http.ResponseWriter, r *http.Request, product ProductStock) {

	sem := make(semaphore, product.Quantity)
	txnStatus := make(chan TransactionData, product.Quantity)
	if remaining := checkStock(product.Product); remaining == 0 {
		goto FINISH
	}

	for i := 0; i < product.Quantity; i++ {
		sem.acquire(1)
		triggerPayment(txnStatus, sem)
	}

	for v := range txnStatus {

		switch v.status {
		case 0:
			sem.acquire(1)
			triggerPayment(txnStatus, sem)
		case 1:
			updateStock(ProductStock{product.Product, -1})
			if remaining := checkStock(product.Product); remaining > 0 {
				sem.acquire(1)
				triggerPayment(txnStatus, sem)
			} else {
				goto FINISH
			}
		}
	}
FINISH:
	json.NewEncoder(w).Encode("All products sold out")

}
