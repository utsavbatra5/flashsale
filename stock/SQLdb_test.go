package stock

import (
	"flashsale/util"
	"reflect"
	"sync"
	"testing"
)

func TestNewSQLdb(t *testing.T) {
	type args struct {
		config util.DBconfig
	}
	tests := []struct {
		name string
		args args
		want Store
	}{
		{"TestNewSQLdb", args{util.DBconfig{
			Key: "",
			URL: "",
		}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSQLdb(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSQLdb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLdb_AddQuantity(t *testing.T) {
	type args struct {
		qtyData qtyData
	}
	tests := []struct {
		name string
		m    *SQLdb
		args args
	}{
		{"TestSQLdb_AddQuantity", &SQLdb{
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

func TestSQLdb_Get(t *testing.T) {
	type args struct {
		product string
	}
	tests := []struct {
		name    string
		m       *SQLdb
		args    args
		want    qtyData
		wantErr bool
	}{
		{"TestSQLdb_Get", &SQLdb{
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
				t.Errorf("SQLdb.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SQLdb.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLdb_UpdateQuantity(t *testing.T) {
	type args struct {
		q qtyData
	}
	tests := []struct {
		name    string
		m       *SQLdb
		args    args
		wantErr bool
	}{
		{"TestSQLdb_UpdateQuantity", &SQLdb{
			mu:         sync.Mutex{},
			totalStock: make(map[string]int),
		}, args{qtyData{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.UpdateQuantity(tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("SQLdb.UpdateQuantity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
