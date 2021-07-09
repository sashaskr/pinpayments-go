package pinpayments

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	var c = http.DefaultClient
	{
		c.Timeout = 25 * time.Second
	}

	tests := []struct{
		name string
		client  *http.Client
	}{
		{
			"nil returns a valid client",
			nil,
		},
		{
			"a passed client is decorated",
			c,
		},
	}

	conf := NewConfig(true, APIEnvSecret, APIEnvKey)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.client, conf)
			if err != nil {
				t.Errorf("not nil error received: %v", err)
			}
		})
	}
}
