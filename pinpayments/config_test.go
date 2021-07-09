package pinpayments

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		t bool
		secretKey string
		publishableKey string
	}
	tests := []struct{
		name string
		args args
		want *Config
	}{
		{
			"config set to testing URL with secret and publishable key",
			args{
				t:              true,
				secretKey:      APIEnvSecret,
				publishableKey: APIEnvKey,
			},
			&Config{
				testing:        true,
				secretKey:      "API_SECRET",
				publishableKey: "API_KEY",
			},
		},
		{
			"config set to live with secret and publishable key",
			args{
				t:              false,
				secretKey:      APIEnvSecret,
				publishableKey: APIEnvKey,
			},
			&Config{
				testing:        false,
				secretKey:      "API_SECRET",
				publishableKey: "API_KEY",
			},
		},
		{
			"config set to live with secret and without publishable",
			args{
				t:              false,
				secretKey:      APIEnvSecret,
				publishableKey: "",
			},
			&Config{
				testing:        false,
				secretKey:      "API_SECRET",
				publishableKey: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.t, tt.args.secretKey, tt.args.publishableKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
