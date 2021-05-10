package main

import (
	"flashsale/addCart"
	"flashsale/stock"
	"flashsale/util"
	"reflect"
	"testing"
)

func Test_initializeStockDB(t *testing.T) {
	type args struct {
		config *util.Config
	}
	tests := []struct {
		name    string
		args    args
		want    stock.Store
		wantErr bool
	}{
		{"Test_initializeStockDB", args{config: &util.Config{Name: "", Stock: []util.ProductStock{}, CartStore: "", CartDBConfig: []util.DBconfig{}, StockStore: "", StockDBConfig: []util.DBconfig{}}}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := initializeStockDB(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("initializeStockDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeStockDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initializeCartDB(t *testing.T) {
	type args struct {
		config *util.Config
	}
	tests := []struct {
		name    string
		args    args
		want    addCart.Store
		wantErr bool
	}{
		{"Test_initializeCartDB", args{config: &util.Config{Name: "", Stock: []util.ProductStock{}, CartStore: "", CartDBConfig: []util.DBconfig{}, StockStore: "", StockDBConfig: []util.DBconfig{}}}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := initializeCartDB(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("initializeCartDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeCartDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
