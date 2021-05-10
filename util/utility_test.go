package util

import (
	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{"TestReadConfig", &Config{
			Name:          "",
			Stock:         nil,
			CartStore:     "",
			CartDBConfig:  nil,
			StockStore:    "",
			StockDBConfig: nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
