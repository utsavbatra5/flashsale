package main

import (
	"context"
	"flashsale/addCart"
	"flashsale/shop"
	"flashsale/stock"
	"flashsale/util"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Initialize Stock Microservice Database instance.
func initializeStockDB(config *util.Config) (stock.Store, error) {
	var dbconfig util.DBconfig
	for _, val := range config.StockDBConfig {
		if val.Key == config.Name {
			dbconfig = val
		}
	}
	return stock.ConfigureDB(config.StockStore, dbconfig)
}

// Initialize Cart Microservice Database instance.
func initializeCartDB(config *util.Config) (addCart.Store, error) {
	var dbconfig util.DBconfig
	for _, val := range config.CartDBConfig {
		if val.Key == config.Name {
			dbconfig = val
		}
	}
	return addCart.ConfigureDB(config.CartStore, dbconfig)
}

// Authentication Middleware
func authMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if user != "user" || pass != "pass" {
			http.Error(w, "Unauthorized.", 401)
			return
		}
		// do stuff
		h.ServeHTTP(w, r)
	})
}

func main() {

	var product shop.ProductStock
	var stockPort, buyPort, addCartPort string

	// Read configurations
	configuration := util.ReadConfig()

	dbInstanceStock, err := initializeStockDB(configuration)
	if err != nil {
		log.Fatalf("failed to instantiate Stock Database", err)
	}

	dbInstanceCart, err := initializeCartDB(configuration)
	if err != nil {
		log.Fatalf("failed to instantiate Stock Database", err)
	}

	// Read Product Name and Stock from config file.
	for _, val := range configuration.Stock {
		product.Quantity = val.Quantity
		product.Product = val.Product
	}

	for _, val := range configuration.Ports {
		switch val.Key {
		case "stock":
			stockPort = val.Port
		case "buy":
			buyPort = val.Port
		case "addCart":
			addCartPort = val.Port
		}
	}

	r1 := mux.NewRouter()
	r2 := mux.NewRouter()
	r3 := mux.NewRouter()

	buyProductHandler := func(w http.ResponseWriter, r *http.Request) {
		shop.BuyProducts(w, r, product)
	}

	getStockHandler := func(w http.ResponseWriter, r *http.Request) {
		stock.GetStock(w, r, dbInstanceStock)
	}
	updateStockHandler := func(w http.ResponseWriter, r *http.Request) {
		stock.UpdateStock(w, r, dbInstanceStock)
	}
	addToCartHandler := func(w http.ResponseWriter, r *http.Request) {
		addCart.AddToCart(w, r, dbInstanceCart)
	}
	getCartDataHandler := func(w http.ResponseWriter, r *http.Request) {
		addCart.GetData(w, r, dbInstanceCart)
	}

	r1 = r1.PathPrefix("/api").Subrouter()
	r1.Use(authMiddleware)
	r2 = r2.PathPrefix("/api").Subrouter()
	r2.Use(authMiddleware)
	r3 = r3.PathPrefix("/api").Subrouter()
	r3.Use(authMiddleware)

	// Router for Stock microservices.
	r1.HandleFunc("/{prodID}", getStockHandler).Methods("GET")
	r1.HandleFunc("/updatestock", updateStockHandler).Methods("POST")

	// Router for Shop/Buy microservice.
	r2.HandleFunc("/buy", buyProductHandler)

	// Router for AddtoCart microservice.
	r3.HandleFunc("/addtocart", addToCartHandler)
	r3.HandleFunc("/getcarts", getCartDataHandler)

	server1 := &http.Server{Addr: ":" + stockPort, Handler: r1}
	go func() {
		if err := server1.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server", err)
		}
	}()

	server2 := &http.Server{Addr: ":" + buyPort, Handler: r2}
	go func() {
		if err := server2.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server", err)
		}
	}()

	server3 := &http.Server{Addr: ":" + addCartPort, Handler: r3}
	go func() {
		if err := server3.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server", err)
		}
	}()

	// Graceful Shutdown of Services
	termChan := make(chan os.Signal)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)

	<-termChan // Blocks here until interrupted

	// Handle shutdown
	fmt.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := server1.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	if err := server2.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	if err := server3.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
}
