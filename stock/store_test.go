package stock

import (
	"net/http"
	"testing"
)

func TestGetStock(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
		s Store
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestGetStock", args{nil, &http.Request{
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
		}, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetStock(tt.args.w, tt.args.r, tt.args.s)
		})
	}
}

func TestUpdateStock(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
		s Store
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestGetStock", args{nil, &http.Request{
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
		}, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateStock(tt.args.w, tt.args.r, tt.args.s)
		})
	}
}
