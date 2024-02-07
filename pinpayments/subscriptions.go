package pinpayments

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type SubscriptionService service

type Subscription struct {
	State                    string    `json:"state,omitempty"`
	NextBillingDate          time.Time `json:"next_billing_date,omitempty"`
	ActiveIntervalStartedAt  time.Time `json:"active_interval_started_at,omitempty"`
	ActiveIntervalFinishedAt time.Time `json:"active_interval_finished_at,omitempty"`
	CancelledAt              time.Time `json:"cancelled_at,omitempty"`
	CreatedAt                time.Time `json:"created_at,omitempty"`
	Token                    string    `json:"token,omitempty"`
	PlanToken                string    `json:"plan_token,omitempty"`
	CustomerToken            string    `json:"customer_token,omitempty"`
	CardToken                string    `json:"card_token,omitempty"`
}

type SubscriptionResponse struct {
	Response Subscription `json:"response,omitempty"`
}

type SubscriptionsResponse struct {
	Response   []Subscription `json:"response,omitempty"`
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

type LedgerResponse struct {
	Response   []Ledger `json:"response,omitempty"`
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

type Ledger struct {
	CreatedAt  time.Time `json:"created_at,omitempty"`
	Type       string    `json:"type,omitempty"`
	Amount     int       `json:"amount,omitempty"`
	Currency   string    `json:"currency,omitempty"`
	Annotation string    `json:"annotation,omitempty"`
}

func (ss *SubscriptionService) Create(subscription *Subscription) (sr *SubscriptionResponse, err error) {
	req, err := ss.client.NewAPIRequest(true, http.MethodPost, "subscriptions", subscription)
	if err != nil {
		return nil, err
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &sr); err != nil {
		return
	}
	return
}

func (ss *SubscriptionService) GetAll(page int) (sr *SubscriptionsResponse, err error) {
	ss.client.SetPage(page)
	req, err := ss.client.NewAPIRequest(true, http.MethodGet, "subscriptions", nil)
	if err != nil {
		return nil, err
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &sr); err != nil {
		return
	}
	return
}

func (ss *SubscriptionService) Get(token string) (sr *SubscriptionResponse, err error) {
	u := fmt.Sprintf("subscriptions/%s", token)
	req, err := ss.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &sr); err != nil {
		return
	}
	return
}

func (ss *SubscriptionService) Update(subscription *Subscription) (sr *SubscriptionResponse, err error) {
	u := fmt.Sprintf("subscriptions/%s", subscription.Token)
	req, err := ss.client.NewAPIRequest(true, http.MethodPut, u, subscription)
	if err != nil {
		return nil, err
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &sr); err != nil {
		return
	}
	return
}

func (ss *SubscriptionService) Delete(token string) (success bool, err error) {
	u := fmt.Sprintf("subscriptions/%s", token)
	req, err := ss.client.NewAPIRequest(true, http.MethodDelete, u, nil)
	if err != nil {
		return false, err
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return false, err
	}
	if res.StatusCode != 204 {
		return false, errors.New("subscription not found")
	}

	return true, nil
}

func (ss *SubscriptionService) ReactivateSubscription(subscription *Subscription) (sr *SubscriptionResponse, err error) {
	u := fmt.Sprintf("subscriptions/%s/reactivate", subscription.Token)
	req, err := ss.client.NewAPIRequest(true, http.MethodPut, u, subscription)
	if err != nil {
		return nil, err
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &sr); err != nil {
		return
	}
	return
}

func (ss *SubscriptionService) GetLedger(token string) (lr *LedgerResponse, err error) {
	u := fmt.Sprintf("subscriptions/%s/ledger", token)
	req, err := ss.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &lr); err != nil {
		return
	}
	return
}
