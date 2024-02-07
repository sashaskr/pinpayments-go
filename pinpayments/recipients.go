package pinpayments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RecipientsService service

type Recipient struct {
	Token            string      `json:"token,omitempty"`
	Email            string      `json:"email,omitempty"`
	Name             string      `json:"name,omitempty"`
	CreatedAt        time.Time   `json:"created_at,omitempty"`
	BankAccount      BankAccount `json:"bank_account,omitempty"`
	BankAccountToken string      `json:"bank_account_token,omitempty"`
}

type RecipientResponse struct {
	Response Recipient `json:"response,omitempty"`
}

type RecipientsResponse struct {
	Response   []Recipient `json:"response,omitempty"`
	Count      int         `json:"count"`
	Pagination struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
		Next     int `json:"next"`
		PerPage  int `json:"per_page"`
		Pages    int `json:"pages"`
		Count    int `json:"count"`
	} `json:"pagination"`
}

func (rs *RecipientsService) Create(recipient *Recipient) (rr *RecipientResponse, err error) {
	req, err := rs.client.NewAPIRequest(true, http.MethodPost, "recipients", recipient)
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

func (rs *RecipientsService) GetAll(page int) (rr *RecipientsResponse, err error) {
	rs.client.SetPage(page)
	req, err := rs.client.NewAPIRequest(true, http.MethodGet, "recipients", nil)
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

func (rs *RecipientsService) Get(token string) (rr *RecipientResponse, err error) {
	u := fmt.Sprintf("recipients/%s", token)
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

func (rs *RecipientsService) Update(recipient *Recipient) (rr *RecipientResponse, err error) {
	u := fmt.Sprintf("recipients/%s", recipient.Token)
	req, err := rs.client.NewAPIRequest(true, http.MethodPut, u, recipient)
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

func (rs *RecipientsService) GetTransfers(token string) (tr *TransfersResponse, err error) {
	u := fmt.Sprintf("recipients/%s/transfers", token)
	req, err := rs.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := rs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &tr); err != nil {
		return
	}
	return
}
