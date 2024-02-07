package pinpayments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RefundsService service

type Refund struct {
	Token         string    `json:"token,omitempty"`
	Success       string    `json:"success,omitempty"`
	Amount        int       `json:"amount,omitempty"`
	Currency      string    `json:"currency,omitempty"`
	Charge        string    `json:"charge,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	ErrorMessage  string    `json:"error_message,omitempty"`
	StatusMessage string    `json:"status_message,omitempty"`
}

type RefundResponse struct {
	Response Refund `json:"response,omitempty"`
}

type RefundsResponse struct {
	Response   []Refund `json:"response,omitempty"`
	Count      int      `json:"count"`
	Pagination struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
		Next     int `json:"next"`
		PerPage  int `json:"per_page"`
		Pages    int `json:"pages"`
		Count    int `json:"count"`
	} `json:"pagination"`
}

type RefundRequest struct {
	ChargeToken string
	Amount      int `json:"amount,omitempty"`
}

func (rs *RefundsService) GetAll(page int) (rr *RefundsResponse, err error) {
	rs.client.SetPage(page)
	req, err := rs.client.NewAPIRequest(true, http.MethodGet, "refunds", nil)
	if err != nil {
		return nil, err
	}

	res, err := rs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &rr); err != nil {
		return
	}
	return
}

func (rs *RefundsService) Get(token string) (rr *RefundResponse, err error) {
	u := fmt.Sprintf("refunds/%s", token)
	req, err := rs.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := rs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &rr); err != nil {
		return
	}
	return
}

func (rs *RefundsService) Create(refundRequest *RefundRequest) (rr *RefundResponse, err error) {
	u := fmt.Sprintf("charges/%s/refunds", refundRequest.ChargeToken)
	req, err := rs.client.NewAPIRequest(true, http.MethodPost, u, refundRequest)
	if err != nil {
		return nil, err
	}

	res, err := rs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &rr); err != nil {
		return
	}
	return
}

func (rs *RefundsService) GetRefundsForCharge(token string, page int) (rr *RefundsResponse, err error) {
	rs.client.SetPage(page)
	u := fmt.Sprintf("charges/%s/refunds", token)
	req, err := rs.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := rs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &rr); err != nil {
		return
	}
	return
}
