package pinpayments

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
	"time"
)

type ChagresService service

type Resp struct {
	Resp interface{} `json:"response"`
}

type ChargesRequest struct {
	Email         string       `json:"email"`
	Description   string       `json:"description"`
	Amount        int32        `json:"amount"`
	IpAddress     string       `json:"ip_address"`
	Currency      string       `json:"currency,omitempty"`
	Capture       bool         `json:"capture,omitempty"`
	Metadata      Metadata     `json:"metadata,omitempty"`
	ThreeDSecure  ThreeDSecure `json:"three_d_secure,omitempty"`
	Card          Card         `json:"card,omitempty"`
	CardToken     string       `json:"card_token,omitempty"`
	CustomerToken string       `json:"customer_token,omitempty"`
}

type ResponseBody struct {
	Token         string    `json:"token,omitempty"`
	Success       bool      `json:"success,omitempty"`
	Amount        int       `json:"amount,omitempty"`
	Currency      string    `json:"currency,omitempty"`
	Description   string    `json:"description,omitempty"`
	Email         string    `json:"email,omitempty"`
	IpAddress     string    `json:"ip_address,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	StatusMessage string    `json:"status_message,omitempty"`
	ErrorMessage  string    `json:"error_message,omitempty"`
	Card          Card      `json:"card,omitempty"`
	Metadata      `json:"metadata,omitempty"`
	TotalFees     int `json:"total_fees,omitempty"`
}

type ChargeResponse struct {
	Response ResponseBody `json:"response"`
}

type ChargesResponse struct {
	Responses  []ResponseBody `json:"response"`
	Count      int            `json:"count"`
	Pagination struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
		Next     int `json:"next"`
		PerPage  int `json:"per_page"`
		Pages    int `json:"pages"`
		Count    int `json:"count"`
	} `json:"pagination"`
}

type Metadata struct {
	OrderNumber        string    `json:"OrderNumber,omitempty"`
	CustomerName       string    `json:"CustomerName,omitempty"`
	OrderTakenBy       string    `json:"order taken by,omitempty"`
	Location           string    `json:"Location,omitempty"`
	TimeOrderCompleted time.Time `json:"time_order_completed,omitempty"`
}

type Card struct {
	Number            string `json:"number,omitempty"`
	ExpiryMonth       int    `json:"expiry_month,omitempty"`
	ExpiryYear        int    `json:"expiry_year,omitempty"`
	CVC               string `json:"cvc,omitempty"`
	Name              string `json:"name,omitempty"`
	AddressLine1      string `json:"address_line1,omitempty"`
	AddressLine2      string `json:"address_line2,omitempty"`
	AddressCity       string `json:"address_city,omitempty"`
	AddressPostcode   string `json:"address_postcode,omitempty"`
	AddressState      string `json:"address_state,omitempty"`
	AddressCountry    string `json:"address_country,omitempty"`
	Scheme            string `json:"scheme,omitempty"`
	DisplayNumber     string `json:"display_number,omitempty"`
	IssuingCountry    string `json:"issuing_country,omitempty"`
	CustomerToken     string `json:"customer_token,omitempty"`
	Token             string `json:"token,omitempty"`
	PublishableApiKey string `json:"publishable_api_key,omitempty"`
}

type ThreeDSecure struct {
	Version       string `json:"version,omitempty"`
	Eci           string `json:"eci,omitempty"`
	Cavv          string `json:"cavv,omitempty"`
	TransactionId string `json:"transaction_id,omitempty"`
}

type Search struct {
	Query     string    `url:"query,omitempty"`
	StartDate time.Time `url:"start_date,omitempty"`
	EndDate   time.Time `url:"end_date,omitempty"`
	Sort      string    `url:"sort,omitempty"`
	Direction int       `url:"direction,omitempty"`
}

func (cs *ChagresService) CreateCharge(charge *ChargesRequest) (cr *ChargeResponse, err error) {
	req, err := cs.client.NewAPIRequest(true, http.MethodPost, "charges", charge)
	if err != nil {
		panic(err)
	}

	res, err := cs.client.Do(req)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(res.content, &cr); err != nil {
		return
	}
	return
}

func (cs *ChagresService) VoidCharge(token string) (cr *ChargeResponse, err error) {
	u := fmt.Sprintf("charges/%s/void", token)
	req, err := cs.client.NewAPIRequest(true, http.MethodPut, u, nil)
	if err != nil {
		panic(err)
	}

	res, err := cs.client.Do(req)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(res.content, &cr); err != nil {
		return
	}
	return
}

func (cs *ChagresService) CaptureCharge(token string) (cr *ChargeResponse, err error) {
	u := fmt.Sprintf("charges/%s/capture", token)
	req, err := cs.client.NewAPIRequest(true, http.MethodPut, u, nil)
	if err != nil {
		panic(err)
	}

	res, err := cs.client.Do(req)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(res.content, &cr); err != nil {
		return
	}
	return
}

func (cs *ChagresService) Get(token string) (cr *ChargeResponse, err error) {
	u := fmt.Sprintf("charges/%s", token)
	req, err := cs.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		panic(err)
	}

	res, err := cs.client.Do(req)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(res.content, &cr); err != nil {
		return
	}
	return
}

func (cs *ChagresService) GetAll() (cr *ChargesResponse, err error) {
	req, err := cs.client.NewAPIRequest(true, http.MethodGet, "charges", nil)
	//TODO: Add pagination here
	if err != nil {
		panic(err)
	}

	res, err := cs.client.Do(req)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(res.content, &cr); err != nil {
		return
	}
	return
}

func (cs *ChagresService) Search(search Search) (cr *ChargesResponse, err error) {
	v, err := query.Values(search)
	if err != nil {
		panic(err)
	}
	u := fmt.Sprintf("charges/search/?%s", v.Encode())
	req, err := cs.client.NewAPIRequest(true, http.MethodGet, u, nil)
	//TODO: Add pagination here
	if err != nil {
		panic(err)
	}

	res, err := cs.client.Do(req)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(res.content, &cr); err != nil {
		return
	}
	return
}
