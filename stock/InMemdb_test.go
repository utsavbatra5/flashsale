package stock

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewInMemdb(t *testing.T) {
	tests := []struct {
		name string
		want Store
	}{
		{"TestNewInMemdb", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemdb(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemdb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemdb_AddQuantity(t *testing.T) {
	type args struct {
		qtyData qtyData
	}
	tests := []struct {
		name string
		m    *InMemdb
		args args
	}{
		{"TestInMemdb_AddQuantity", &InMemdb{
			mu:         sync.Mutex{},
			totalStock: make(map[string]int),
		}, args{qtyData{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.AddQuantity(tt.args.qtyData)
		})
	}
}

func TestInMemdb_Get(t *testing.T) {
	type args struct {
		product string
	}
	tests := []struct {
		name    string
		m       *InMemdb
		args    args
		want    qtyData
		wantErr bool
	}{
		{"TestInMemdb_Get", &InMemdb{
			mu:         sync.Mutex{},
			totalStock: make(map[string]int),
		}, args{""}, qtyData{
			Product:  "",
			Quantity: 0,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Get(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemdb.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemdb.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemdb_UpdateQuantity(t *testing.T) {
	type args struct {
		q qtyData
	}
	tests := []struct {
		name    string
		m       *InMemdb
		args    args
		wantErr bool
	}{
		{"TestInMemdb_UpdateQuantity", &InMemdb{
			mu:         sync.Mutex{},
			totalStock: make(map[string]int),
		}, args{qtyData{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.UpdateQuantity(tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("InMemdb.UpdateQuantity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
