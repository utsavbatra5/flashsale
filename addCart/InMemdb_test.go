package addCart

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

func TestGetInMem(t *testing.T) {
	tests := []struct {
		name string
		want Store
	}{
		{"TestNewInMemdb", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInMem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInMem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMem_AddQuantity(t *testing.T) {
	type args struct {
		user     int64
		product  string
		quantity int
	}
	tests := []struct {
		name string
		m    *InMem
		args args
	}{
		{"TestInMemdb_AddQuantity", &InMem{
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

func TestInMem_Get(t *testing.T) {
	tests := []struct {
		name    string
		m       *InMem
		want    []cart
		wantErr bool
	}{
		{"TestInMem_Get", &InMem{
			mu:           sync.Mutex{},
			productStock: []cart{},
		}, []cart{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("InMem.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMem.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMem_Add(t *testing.T) {
	type args struct {
		c cart
	}
	tests := []struct {
		name    string
		m       *InMem
		args    args
		wantErr bool
	}{
		{"TestInMem_Add", &InMem{
			mu:           sync.Mutex{},
			productStock: []cart{},
		}, args{cart{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Add(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("InMem.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
