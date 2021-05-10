package addCart

import (
	"flashsale/util"
	"reflect"
	"testing"
)

func TestConfigureDB(t *testing.T) {
	type args struct {
		db       string
		dbconfig util.DBconfig
	}
	tests := []struct {
		name    string
		args    args
		want    Store
		wantErr bool
	}{
		{"TestConfigureDB", args{"", util.DBconfig{"", ""}}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConfigureDB(tt.args.db, tt.args.dbconfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigureDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigureDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
