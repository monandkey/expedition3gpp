package expedition

import (
	"reflect"
	"testing"
)

func Test_configLoad(t *testing.T) {
	tests := []struct {
		name string
		want configParams
	}{
		{
			name: "Normal Test",
			want: configParams{
				strageLocation:     "HOMEDIR",
				cacheEnable:        true,
				cacheRetentionTime: 14400,
				cacheLocation:      "HOMEDIR",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configLoad(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configLoad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCacheFile(t *testing.T) {
	type args struct {
		cacheLocation  string
		documentNumber string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal Test HOMEDIR",
			args: args{
				cacheLocation:  "HOMEDIR",
				documentNumber: "23.401",
			},
			want: "/root/.cache/23401.yaml",
		},
		{
			name: "Normal Test /root",
			args: args{
				cacheLocation:  "/root",
				documentNumber: "23.401",
			},
			want: "/root/.cache/23401.yaml",
		},
		{
			name: "Normal Test /root/",
			args: args{
				cacheLocation:  "/root/",
				documentNumber: "23.401",
			},
			want: "/root/.cache/23401.yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCacheFileName(tt.args.cacheLocation, tt.args.documentNumber); got != tt.want {
				t.Errorf("getCacheFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cacheValidate(t *testing.T) {
	type args struct {
		cacheRetentionTime int
		fileName           string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Normal Test",
			args: args{
				cacheRetentionTime: 14400,
				fileName:           "/root/.cache/23401.yaml",
			},
			want: true,
		},
		{
			name: "Normal Test file not found",
			args: args{
				cacheRetentionTime: 14400,
				fileName:           "/root.cache/23401.yaml",
			},
			want: false,
		},
		{
			name: "Normal Test cache time over",
			args: args{
				cacheRetentionTime: 1,
				fileName:           "/root/.cache/23401.yaml",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cacheValidate(tt.args.cacheRetentionTime, tt.args.fileName); got != tt.want {
				t.Errorf("cacheValidate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cacheLoad(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Normal Test",
			args: args{
				filePath: "/root/.cache/23401.yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := cacheLoad(tt.args.filePath)
			formatDisplayAll(value.Value)
		})
	}
}
