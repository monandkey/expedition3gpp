package config

import (
	"testing"
)

func Test_baseParams_Load(t *testing.T) {
	type fields struct {
		params params
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Normal Test",
			fields: fields{
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
			b := baseParams{
				params: tt.fields.params,
			}
			b.Load()
		})
	}
}

func Test_baseParams_Write(t *testing.T) {
	type fields struct {
		params params
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Normal Test",
			fields: fields{
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
			b := baseParams{
				params: tt.fields.params,
			}
			b.Write()
		})
	}
}
