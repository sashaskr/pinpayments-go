package pinpayments

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
	"time"
)

type TransfersService service

type Transfer struct {
	Token          string      `json:"token,omitempty"`
	Status         string      `json:"status,omitempty"`
	Currency       string      `json:"currency,omitempty"`
	Description    string      `json:"description,omitempty"`
	Amount         int         `json:"amount,omitempty"`
	TotalDebits    int         `json:"total_debits,omitempty"`
	TotalCredits   int         `json:"total_credits,omitempty"`
	CreatedAt      time.Time   `json:"created_at,omitempty"`
	PaidAt         time.Time   `json:"paid_at,omitempty"`
	LineItemsCount int         `json:"line_items_count,omitempty"`
	BankAccount    BankAccount `json:"bank_account,omitempty"`
	Recipient      string      `json:"recipient,omitempty"`
}

type TransferResponse struct {
	Response Transfer `json:"response,omitempty"`
}

type TransfersResponse struct {
	Response   []Transfer `json:"response,omitempty"`
	Count      int        `json:"count"`
	Pagination struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
		Next     int `json:"next"`
		PerPage  int `json:"per_page"`
		Pages    int `json:"pages"`
		Count    int `json:"count"`
	} `json:"pagination"`
}

type LineItem struct {
	Type      string    `json:"type,omitempty"`
	Amount    int       `json:"amount,omitempty"`
	Currency  string    `json:"currency,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Object    string    `json:"object,omitempty"`
	Token     string    `json:"token,omitempty"`
	Record    Record    `json:"record,omitempty"`
}

type Record struct {
	Type      string    `json:"type,omitempty"`
	Amount    int       `json:"amount,omitempty"`
	Currency  string    `json:"currency,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type LineItemsResponse struct {
	Response   []LineItem `json:"response,omitempty"`
	Count      int        `json:"count"`
	Pagination struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
		Next     int `json:"next"`
		PerPage  int `json:"per_page"`
		Pages    int `json:"pages"`
		Count    int `json:"count"`
	} `json:"pagination"`
}

func (ts *TransfersService) Create(transfer *Transfer) (tr *TransferResponse, err error) {
	req, err := ts.client.NewAPIRequest(true, http.MethodPost, "transfers", transfer)
	if err != nil {
		return nil, err
	}

	res, err := ts.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &tr); err != nil {
		return
	}
	return
}

func (ts *TransfersService) GetAll(page int) (tr *TransfersResponse, err error) {
	ts.client.SetPage(page)
	req, err := ts.client.NewAPIRequest(true, http.MethodGet, "transfers", nil)
	if err != nil {
		return nil, err
	}

	res, err := ts.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &tr); err != nil {
		return
	}
	return
}

func (ts *TransfersService) Get(token string) (tr *TransferResponse, err error) {
	u := fmt.Sprintf("transfers/%s", token)
	req, err := ts.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := ts.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &tr); err != nil {
		return
	}
	return
}

func (ts *TransfersService) Search(search Search) (tr *TransfersResponse, err error) {
	ts.client.SetPage(search.Page)
	v, err := query.Values(search)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("transfers/search/?%s", v.Encode())
	req, err := ts.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := ts.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &tr); err != nil {
		return
	}
	return
}

func (ts *TransfersService) GetLineItems(token string, page int) (lr *LineItemsResponse, err error) {
	ts.client.SetPage(page)
	u := fmt.Sprintf("transfers/%s/line_items", token)
	req, err := ts.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := ts.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &lr); err != nil {
		return
	}
	return
}
