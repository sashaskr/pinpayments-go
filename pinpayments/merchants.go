package pinpayments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type MerchantsService service

type Merchant struct {
	Contact     Contact     `json:"contact,omitempty"`
	Entity      Entity      `json:"entity,omitempty"`
	Business    Business    `json:"business,omitempty"`
	BankAccount BankAccount `json:"bank_account,omitempty"`
	Director    Director    `json:"director,omitempty"`
	Notes       string      `json:"notes,omitempty"`
}

type Contact struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
}

type Entity struct {
	BusinessRegistrationNumber string `json:"business_registration_number,omitempty"`
	FullLegalName              string `json:"full_legal_name,omitempty"`
	IncorporationStatus        string `json:"incorporation_status,omitempty"`
	RegisteredForGst           bool   `json:"registered_for_gst,omitempty"`
	AddressLine1               string `json:"address_line_1,omitempty"`
	AddressLine2               string `json:"address_line_2,omitempty"`
	AddressLocality            string `json:"address_locality,omitempty"`
	AddressRegion              string `json:"address_region,omitempty"`
	AddressPostalCode          string `json:"address_postal_code,omitempty"`
}

type Business struct {
	TradingName             string `json:"trading_name,omitempty"`
	Description             string `json:"description,omitempty"`
	TypicalProductPrice     string `json:"typical_product_price,omitempty"`
	TransactionsPerMonth    string `json:"transactions_per_month,omitempty"`
	AnnualTransactionVolume string `json:"annual_transaction_volume,omitempty"`
	SellsPhysicalGoods      string `json:"sells_physical_goods,omitempty"`
	AverageDeliveryDays     string `json:"average_delivery_days,omitempty"`
	URL                     string `json:"url,omitempty"`
}

type Director struct {
	FullName      string `json:"full_name,omitempty"`
	ContactNumber string `json:"contact_number,omitempty"`
	DateOfBirth   string `json:"date_of_birth,omitempty"`
}

type MerchantResponse struct {
	Response struct {
		Token        string    `json:"token,omitempty"`
		Email        string    `json:"email,omitempty"`
		BusinessName string    `json:"business_name,omitempty"`
		Status       string    `json:"status,omitempty"`
		CreatedAt    time.Time `json:"created_at,omitempty"`
		UpdatedAt    time.Time `json:"updated_at,omitempty"`
	} `json:"response,omitempty"`
}

type MerchantsResponse struct {
	Response []struct {
		Token        string    `json:"token,omitempty"`
		Email        string    `json:"email,omitempty"`
		BusinessName string    `json:"business_name,omitempty"`
		Status       string    `json:"status,omitempty"`
		CreatedAt    time.Time `json:"created_at,omitempty"`
		UpdatedAt    time.Time `json:"updated_at,omitempty"`
	} `json:"response,omitempty"`
	Pagination struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
		Next     int `json:"next"`
		PerPage  int `json:"per_page"`
		Pages    int `json:"pages"`
		Count    int `json:"count"`
	} `json:"pagination"`
}

type MerchantFullResponse struct {
	Response struct {
		Token                            string    `json:"token,omitempty"`
		Status                           string    `json:"status,omitempty"`
		BusinessName                     string    `json:"business_name,omitempty"`
		ContactName                      string    `json:"contact_name,omitempty"`
		ContactEmail                     string    `json:"contact_email,omitempty"`
		SkipChargeAddressValidation      bool      `json:"skip_charge_address_validation,omitempty"`
		SendTransferEmail                bool      `json:"send_transfer_email,omitempty"`
		SendRefundSuccessfulEmails       bool      `json:"send_refund_successful_emails,omitempty"`
		SettlementDelayDays              int       `json:"settlement_delay_days,omitempty"`
		AmexEnabled                      bool      `json:"amex_enabled,omitempty"`
		MaximumChargeCents               int       `json:"maximum_charge_cents,omitempty"`
		PosMaximumChargeCents            int       `json:"pos_maximum_charge_cents,omitempty"`
		TransfersEnabled                 bool      `json:"transfers_enabled,omitempty"`
		LastChargeDate                   time.Time `json:"last_charge_date,omitempty"`
		TestSecretApiKey                 string    `json:"test_secret_api_key,omitempty"`
		TestPublishableApiKey            string    `json:"test_publishable_api_key,omitempty"`
		TransactionFixedFee              int       `json:"transaction_fixed_fee,omitempty"`
		ThirdPartyTransferFee            int       `json:"third_party_transfer_fee,omitempty"`
		DomesticTransactionPercentageFee float64   `json:"domestic_transaction_percentage_fee,omitempty"`
		AmexTransactionPercentageFee     float64   `json:"amex_transaction_percentage_fee,omitempty"`
		ForeignTransactionPercentageFee  float64   `json:"foreign_transaction_percentage_fee,omitempty"`
		Amex                             bool      `json:"amex,omitempty"`
		TransferEmailLineItemsAttachment bool      `json:"transfer_email_line_items_attachment,omitempty"`
		StatementDescriptorRandomCode    bool      `json:"statement_descriptor_random_code,omitempty"`
		ApiKeys                          struct {
			Secret      string `json:"secret,omitempty"`
			Publishable string `json:"publishable,omitempty"`
		} `json:"api_keys,omitempty"`
		Pricing struct {
			DomesticCurrencyDomesticCardCore    Fee `json:"domestic_currency_domestic_card_core,omitempty"`
			DomesticCurrencyDomesticCardPremium Fee `json:"domestic_currency_domestic_card_premium,omitempty"`
			DomesticCurrencyForeignCardCore     Fee `json:"domestic_currency_foreign_card_core,omitempty"`
			DomesticCurrencyForeignCardPremium  Fee `json:"domestic_currency_foreign_card_premium,omitempty"`
			ForeignCurrencyCore                 Fee `json:"foreign_currency_core,omitempty"`
			ForeignCurrencyPremium              Fee `json:"foreign_currency_premium,omitempty"`
		} `json:"pricing,omitempty"`
		Notifications struct {
			Receipts struct {
				Api               bool `json:"api,omitempty"`
				PaymentPage       bool `json:"payment_page,omitempty"`
				ManualCharge      bool `json:"manual_charge,omitempty"`
				RecurringPayments bool `json:"recurring_payments,omitempty"`
				XeroPaymentPage   bool `json:"xero_payment_page,omitempty"`
			} `json:"receipts,omitempty"`
			RecurringPayments struct {
				Subscribed    bool `json:"subscribed,omitempty"`
				Unsubscribed  bool `json:"unsubscribed,omitempty"`
				Cancelled     bool `json:"cancelled,omitempty"`
				Renewed       bool `json:"renewed,omitempty"`
				RenewalFailed bool `json:"renewal_failed,omitempty"`
				Reactivated   bool `json:"reactivated,omitempty"`
			} `json:"recurring_payments,omitempty"`
		} `json:"notifications,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	} `json:"response,omitempty"`
}

type Fee struct {
	FixedFee      int     `json:"fixed_fee,omitempty"`
	PercentageFee float64 `json:"percentage_fee,omitempty"`
}

func (ms *MerchantsService) Create(merchant *Merchant) (mr *MerchantResponse, err error) {
	req, err := ms.client.NewAPIRequest(true, http.MethodPost, "merchants", merchant)
	if err != nil {
		return nil, err
	}

	res, err := ms.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &mr); err != nil {
		return
	}
	return
}

func (ms *MerchantsService) GetAll(page int) (mr *MerchantsResponse, err error) {
	ms.client.SetPage(page)
	req, err := ms.client.NewAPIRequest(true, http.MethodGet, "merchants", nil)
	if err != nil {
		return nil, err
	}

	res, err := ms.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &mr); err != nil {
		return
	}
	return
}

func (ms *MerchantsService) Get(token string) (mr *MerchantFullResponse, err error) {
	u := fmt.Sprintf("merchants/%s", token)
	req, err := ms.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := ms.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &mr); err != nil {
		return
	}
	return
}

func (ms *MerchantsService) GetDefault(token string) (mr *MerchantFullResponse, err error) {
	req, err := ms.client.NewAPIRequest(true, http.MethodGet, "merchants/default_settings", nil)
	if err != nil {
		return nil, err
	}

	res, err := ms.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &mr); err != nil {
		return
	}
	return
}
