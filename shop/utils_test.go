package shop

import (
	"reflect"
	"testing"
)

func Test_semaphore_acquire(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		sem  semaphore
		args args
	}{
		{"Test_semaphore_acquire", nil, args{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sem.acquire(tt.args.n)
		})
	}
}

func Test_semaphore_release(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		sem  semaphore
		args args
	}{
		{"Test_semaphore_release", nil, args{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sem.release(tt.args.n)
		})
	}
}

func Test_getVal(t *testing.T) {
	tests := []struct {
		name    string
		want    TransactionData
		wantErr bool
	}{
		{"Test_getVal", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getVal()
			if (err != nil) != tt.wantErr {
				t.Errorf("getVal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executePayment(t *testing.T) {
	type args struct {
		d   TransactionData
		val chan TransactionData
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test_executePayment", args{nil, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executePayment(tt.args.d, tt.args.val)
		})
	}
}
