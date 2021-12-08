package expedition

import (
	"testing"
)

func Test_getSaveFilePath(t *testing.T) {
	type args struct {
		savePath       string
		documentNumber string
		url            string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal Test",
			args: args{
				savePath:       "HOMEDIR",
				documentNumber: "23401",
				url:            "https://www.3gpp.org/ftp/Specs/archive/23_series/23.401/23401-110.zip",
			},
			want: "/root/23401-110.zip",
		},
		{
			name: "not HOMEDIR Test",
			args: args{
				savePath:       "/root",
				documentNumber: "23401",
				url:            "https://www.3gpp.org/ftp/Specs/archive/23_series/23.401/23401-110.zip",
			},
			want: "/root/23401-110.zip",
		},
		{
			name: "not HOMEDIR and separate Test",
			args: args{
				savePath:       "/root/",
				documentNumber: "23401",
				url:            "https://www.3gpp.org/ftp/Specs/archive/23_series/23.401/23401-110.zip",
			},
			want: "/root/23401-110.zip",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSaveFilePath(tt.args.savePath, tt.args.documentNumber, tt.args.url); got != tt.want {
				t.Errorf("getSaveFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileUnzip(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Normal Test",
			args: args{
				path: "/root/23401-h20.zip",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := fileUnzip(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("fileUnzip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
