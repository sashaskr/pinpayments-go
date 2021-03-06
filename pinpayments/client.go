package pinpayments

import (
	"net/http"
	"net/url"
	"runtime"
	"strings"
)

const (
	BaseURL            string = "https://api.pinpayments.com/1/"
	BaseURLTest        string = "https://test-api.pinpayments.com/1/"
	RequestContentType string = "application/json"
	RequestAccept      string = "application/json"
	Connection         string = "keep-alive"
	AuthHeader         string = "Authorization"
	TokenType          string = "Basic"
	APIEnvSecret       string = "API_SECRET"
	APIEnvKey          string = "API_KEY"
)

type Client struct {
	PaginationInterface
	page             *int
	BaseURL          *url.URL
	config           *Config
	userAgent        string
	client           *http.Client
	common           service
	secretKey        string
	publishableKey   string
	Charges          *ChargesService
	Customers        *CustomersService
	Refunds          *RefundsService
	Cards            *CardsService
	Recipients       *RecipientsService
	Transfers        *TransfersService
	Balance          *BalanceService
	BankAccount      *BankAccountService
	Events           *EventsService
	WebhookEndpoints *WebhookEndpointsService
	Webhooks         *WebhooksService
	Plans            *PlansService
	Subscription     *SubscriptionService
	Merchant         *MerchantsService
}

type service struct {
	client *Client
}

type PaginationInterface interface {
	SetPage(page int)
}

func (c *Client) SetPage(page int) {
	c.page = &page
}

func (c *Client) WithAuthenticationValue(k string, p string) error {
	if k == "" {
		return errEmptyApiKey
	}

	if p == "" {
		return errEmptyApiPublishable
	}

	c.secretKey = strings.TrimSpace(k)
	c.publishableKey = strings.TrimSpace(p)
	return nil
}

func NewClient(baseClient *http.Client, c *Config) (pinpayments *Client, err error) {
	if baseClient == nil {
		baseClient = http.DefaultClient
	}
	var host string

	if c.testing {
		host = BaseURLTest
	} else {
		host = BaseURL
	}

	u, _ := url.Parse(host)

	pinpayments = &Client{
		BaseURL: u,
		config:  c,
		client:  baseClient,
	}

	pinpayments.common.client = pinpayments

	// here is all services begin
	pinpayments.Charges = (*ChargesService)(&pinpayments.common)
	pinpayments.Customers = (*CustomersService)(&pinpayments.common)
	pinpayments.Refunds = (*RefundsService)(&pinpayments.common)
	pinpayments.Cards = (*CardsService)(&pinpayments.common)
	pinpayments.Recipients = (*RecipientsService)(&pinpayments.common)
	pinpayments.Transfers = (*TransfersService)(&pinpayments.common)
	pinpayments.Balance = (*BalanceService)(&pinpayments.common)
	pinpayments.BankAccount = (*BankAccountService)(&pinpayments.common)
	pinpayments.Events = (*EventsService)(&pinpayments.common)
	pinpayments.WebhookEndpoints = (*WebhookEndpointsService)(&pinpayments.common)
	pinpayments.Webhooks = (*WebhooksService)(&pinpayments.common)
	pinpayments.Plans = (*PlansService)(&pinpayments.common)
	pinpayments.Subscription = (*SubscriptionService)(&pinpayments.common)
	pinpayments.Merchant = (*MerchantsService)(&pinpayments.common)
	// here is all services end

	pinpayments.publishableKey = c.publishableKey
	pinpayments.secretKey = c.secretKey
	pinpayments.userAgent = strings.Join([]string{
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Version(),
	}, ";")
	return
}
