package daemon

import (
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name:    "Parse config file in yaml format",
			args:    args{filename: "../../build/config.yaml"},
			want:    Config{Database: "/var/lib/reef-pi/reef-pi.db"},
			wantErr: false,
		},
		{
			name:    "Parse config file in json format",
			args:    args{filename: "../../build/config.json"},
			want:    Config{Database: "/var/lib/reef-pi/reef-pi.db"},
			wantErr: false,
		},
		{
			name:    "Default config on error",
			args:    args{filename: "../../build/config.jpeg"},
			want:    Config{Database: "reef-pi.db"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseConfig(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
