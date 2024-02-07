package pinpayments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type ChargesService service

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
	Reference     string       `json:"reference,omitempty"`
	Metadata      Metadata     `json:"metadata,omitempty"`
	ThreeDSecure  ThreeDSecure `json:"three_d_secure,omitempty"`
	Card          *Card        `json:"card,omitempty"`
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
	OrderNumber        string     `json:"order_number,omitempty"`
	CustomerName       string     `json:"customer_name,omitempty"`
	OrderTakenBy       string     `json:"order_taken_by,omitempty"`
	Location           string     `json:"location,omitempty"`
	TimeOrderCompleted *time.Time `json:"time_order_completed,omitempty"`
}

type ThreeDSecure struct {
	Version       string `json:"version,omitempty"`
	Eci           string `json:"eci,omitempty"`
	Cavv          string `json:"cavv,omitempty"`
	TransactionId string `json:"transaction_id,omitempty"`
}

type Search struct {
	Page      int
	Query     string    `url:"query,omitempty"`
	StartDate time.Time `url:"start_date,omitempty"`
	EndDate   time.Time `url:"end_date,omitempty"`
	Sort      string    `url:"sort,omitempty"`
	Direction int       `url:"direction,omitempty"`
}

func (cs *ChargesService) CreateCharge(charge *ChargesRequest) (cr *ChargeResponse, err error) {
	req, err := cs.client.NewAPIRequest(true, http.MethodPost, "charges", charge)
	if err != nil {
		return nil, err
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &cr); err != nil {
		return
	}
	return
}

func (cs *ChargesService) VoidCharge(token string) (cr *ChargeResponse, err error) {
	u := fmt.Sprintf("charges/%s/void", token)
	req, err := cs.client.NewAPIRequest(true, http.MethodPut, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &cr); err != nil {
		return
	}
	return
}

func (cs *ChargesService) CaptureCharge(token string) (cr *ChargeResponse, err error) {
	u := fmt.Sprintf("charges/%s/capture", token)
	req, err := cs.client.NewAPIRequest(true, http.MethodPut, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &cr); err != nil {
		return
	}
	return
}

func (cs *ChargesService) Get(token string) (cr *ChargeResponse, err error) {
	u := fmt.Sprintf("charges/%s", token)
	req, err := cs.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &cr); err != nil {
		return
	}
	return
}

func (cs *ChargesService) GetAll(page int) (cr *ChargesResponse, err error) {
	cs.client.SetPage(page)
	req, err := cs.client.NewAPIRequest(true, http.MethodGet, "charges", nil)
	if err != nil {
		return nil, err
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &cr); err != nil {
		return
	}
	return
}

func (cs *ChargesService) Search(search Search) (cr *ChargesResponse, err error) {
	cs.client.SetPage(search.Page)
	v, err := query.Values(search)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("charges/search/?%s", v.Encode())
	req, err := cs.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &cr); err != nil {
		return
	}
	return
}
