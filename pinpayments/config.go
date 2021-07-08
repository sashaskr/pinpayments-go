package pinpayments

type Config struct {
	testing        bool
	secretKey      string
	publishableKey string
}

func NewConfig(t bool, secretKey string, publishableKey string) *Config {
	return &Config{
		testing:        t,
		secretKey:      secretKey,
		publishableKey: publishableKey,
	}
}
