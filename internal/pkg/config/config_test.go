package config

import (
	"reflect"
	"testing"
)

func Test_getFileName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "normal test",
			want: "/root/.expedition3gpp.yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFileName(); got != tt.want {
				t.Errorf("getFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileExist(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Normal test",
			args: args{
				fileName: "/root/.expedition3gpp.yaml",
			},
			want: true,
		},
		{
			name: "Failure test",
			args: args{
				fileName: "/home/.expedition3gpp.yaml",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileExist(tt.args.fileName); got != tt.want {
				t.Errorf("fileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_configLoad(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want params
	}{
		{
			name: "Normal test",
			args: args{
				fileName: "/root/.expedition3gpp.yaml",
			},
			want: params{
				StrageLocation:     "HOMEDIR",
				CacheEnable:        true,
				CacheRetentionTime: 14400,
				CacheLocation:      "HOMEDIR",
			},
		},
		{
			name: "Failure test",
			args: args{
				fileName: "/home/.expedition3gpp.yaml",
			},
			want: params{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := configLoad(tt.args.fileName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configLoad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_configWrite(t *testing.T) {
	type args struct {
		fileName string
		params   params
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Normal tst",
			args: args{
				fileName: "/root/.expedition3gpp.yaml",
				params: params{
					StrageLocation:     "HOMEDIR",
					CacheEnable:        true,
					CacheRetentionTime: 14400,
					CacheLocation:      "HOMEDIR",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := configWrite(tt.args.fileName, tt.args.params); err != nil {
				t.Errorf("%s", err)
			}
		})
	}
}
