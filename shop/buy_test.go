package shop

import (
	"net/http"
	"testing"
)

func Test_checkStock(t *testing.T) {
	type args struct {
		product string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_checkStock", args{""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStock(tt.args.product); got != tt.want {
				t.Errorf("checkStock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateStock(t *testing.T) {
	type args struct {
		ps ProductStock
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test_updateStock", args{ProductStock{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := updateStock(tt.args.ps); (err != nil) != tt.wantErr {
				t.Errorf("updateStock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_triggerPayment(t *testing.T) {
	type args struct {
		ch  chan TransactionData
		sem semaphore
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test_triggerPayment", args{nil, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			triggerPayment(tt.args.ch, tt.args.sem)
		})
	}
}

func TestBuyProducts(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		r       *http.Request
		product ProductStock
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestBuyProducts", args{nil, &http.Request{
			Method:           "",
			URL:              nil,
			Proto:            "",
			ProtoMajor:       0,
			ProtoMinor:       0,
			Header:           nil,
			Body:             nil,
			GetBody:          nil,
			ContentLength:    0,
			TransferEncoding: nil,
			Close:            false,
			Host:             "",
			Form:             nil,
			PostForm:         nil,
			MultipartForm:    nil,
			Trailer:          nil,
			RemoteAddr:       "",
			RequestURI:       "",
			TLS:              nil,
			Cancel:           nil,
			Response:         nil,
		}, ProductStock{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BuyProducts(tt.args.w, tt.args.r, tt.args.product)
		})
	}
}
