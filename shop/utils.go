package shop

import (
	"math/rand"
	"time"
)

type TransactionData struct {
	orderID int
	userID  int
	status  int
}
type empty struct{}

type semaphore chan empty

// acquire n resources
func (sem semaphore) acquire(n int) {
	e := empty{}
	for i := 0; i < n; i++ {
		sem <- e
	}
}

// release n resources
func (sem semaphore) release(n int) {
	for i := 0; i < n; i++ {
		<-sem
	}
}

// generator to generate random User Cart data for simulation
func getVal() (TransactionData, error) {
	orderID := rand.Intn(1000000)
	return TransactionData{orderID, 1, -1}, nil
}

// dummy mechanism to execute payment.
func executePayment(d TransactionData, val chan TransactionData) {

	c1 := make(chan int, 1)
	go func() {
		c1 <- 1
	}()

	select {
	// timeout for payment completion.
	case <-time.After(180 * time.Second):
		d.status = 1
		val <- d

	case <-c1:
		x := rand.Intn(10000)
		d.userID = x
		// dummy logic generator for payment status. If condition passes,
		// payment is successful. Otherwise it fails.
		if x%5 == 0 && x%9 == 0 {
			d.status = 1
			val <- d
		} else {
			d.status = 0
			val <- d
		}
	}

}
