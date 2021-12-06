package expedition

import (
	"errors"
	"testing"
)

func Test_fetchPage(t *testing.T) {
	type want struct {
		url     string
		version string
	}
	type args struct {
		docNum string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Normal Test",
			args: args{
				docNum: "23401",
			},
			want: want{
				url:     "https://www.3gpp.org/ftp/Specs/archive/23_series/23.401/23401-h20.zip",
				version: "17.0.0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pageFetch(tt.args.docNum)
			if err != nil {
				t.Errorf("Err: %s", err)
			}
			if got[0].Url != tt.want.url && got[0].Version != tt.want.version {
				t.Errorf("URL: %s version: %s", got[0].Url, got[0].Version)
			}
		})
	}
}

func Test_downloadContents(t *testing.T) {
	type args struct {
		url      string
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Normal Test",
			args: args{
				url:      "https://www.3gpp.org/ftp/Specs/archive/23_series/23.401/23401-h20.zip",
				fileName: "/root/23401-h20.zip",
			},
			wantErr: nil,
		},
		{
			name: "Failure Test",
			args: args{
				url:      "https://www.3gpp.org/ftp/Specs/archive/23_series/23.401/23401-h20",
				fileName: "/root/23401-h20.zip",
			},
			wantErr: ErrCode403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := downloadContents(tt.args.url, tt.args.fileName)
			if !(errors.Is(err, tt.wantErr)) {
				t.Errorf("downloadContents() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getDownloadURL(t *testing.T) {
	type args struct {
		valu    []valueBody
		version string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal Test",
			args: args{
				valu: []valueBody{
					{
						Version: "1.0.0",
						Name:    "23.401",
						Url:     "https://www.3gpp.org/ftp/Specs/archive/23_series/23.401/23401-110.zip",
					},
					{
						Version: "1.0.1",
						Name:    "23.401",
						Url:     "https://www.3gpp.org/ftp/Specs/archive/23_series/23.401/23401-110.zip",
					},
				},
				version: "1.0.0",
			},
			want: "https://www.3gpp.org/ftp/Specs/archive/23_series/23.401/23401-110.zip",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDownloadURL(tt.args.valu, tt.args.version); got != tt.want {
				t.Errorf("getDownloadURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
