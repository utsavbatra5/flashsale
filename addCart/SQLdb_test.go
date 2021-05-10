package addCart

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
		{"TestNewSQLdb", args{util.DBconfig{}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSQLdb(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSQLdb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSQLdb(t *testing.T) {
	tests := []struct {
		name string
		want Store
	}{
		{"TestGetSQLdb", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSQLdb(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSQLdb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLdb_AddQuantity(t *testing.T) {
	type args struct {
		user     int64
		product  string
		quantity int
	}
	tests := []struct {
		name string
		m    *SQLdb
		args args
	}{
		{"TestSQLdb_AddQuantity", &SQLdb{
			mu:           sync.Mutex{},
			productStock: []cart{},
		}, args{0, "", 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.AddQuantity(tt.args.user, tt.args.product, tt.args.quantity)
		})
	}
}

func TestSQLdb_Get(t *testing.T) {
	tests := []struct {
		name    string
		m       *SQLdb
		want    []cart
		wantErr bool
	}{
		{"TestSQLdb_Get", &SQLdb{
			mu:           sync.Mutex{},
			productStock: []cart{},
		}, []cart{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Get()
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

func TestSQLdb_Add(t *testing.T) {
	type args struct {
		c cart
	}
	tests := []struct {
		name    string
		m       *SQLdb
		args    args
		wantErr bool
	}{
		{"TestSQLdb_Add", &SQLdb{
			mu:           sync.Mutex{},
			productStock: []cart{},
		}, args{cart{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Add(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SQLdb.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
