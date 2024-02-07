package pinpayments

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type PlansService service

type Plan struct {
	Name                string             `json:"name,omitempty"`
	Amount              int                `json:"amount,omitempty"`
	Currency            string             `json:"currency,omitempty"`
	SetupAmount         int                `json:"setup_amount,omitempty"`
	TrialAmount         int                `json:"trial_amount,omitempty"`
	Interval            int                `json:"interval,omitempty"`
	IntervalUnit        string             `json:"interval_unit,omitempty"`
	Intervals           int                `json:"intervals,omitempty"`
	TrialInterval       int                `json:"trial_interval,omitempty"`
	TrialIntervalUnit   string             `json:"trial_interval_unit,omitempty"`
	CreatedAt           time.Time          `json:"created_at,omitempty"`
	Token               string             `json:"token,omitempty"`
	CustomerPermissions []string           `json:"customer_permissions,omitempty"`
	SubscriptionCounts  SubscriptionCounts `json:"subscription_counts,omitempty"`
}

type SubscriptionCounts struct {
	Trial      int `json:"trial,omitempty"`
	Active     int `json:"active,omitempty"`
	Cancelling int `json:"cancelling,omitempty"`
	Cancelled  int `json:"cancelled,omitempty"`
}

type PlanResponse struct {
	Response Plan `json:"response,omitempty"`
}

type PlansResponse struct {
	Response   []Plan `json:"response,omitempty"`
	Pagination struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
		Next     int `json:"next"`
		PerPage  int `json:"per_page"`
		Pages    int `json:"pages"`
		Count    int `json:"count"`
	} `json:"pagination"`
}

func (ps *PlansService) Create(plan *Plan) (pr *PlanResponse, err error) {
	req, err := ps.client.NewAPIRequest(true, http.MethodPost, "plans", plan)
	if err != nil {
		return nil, err
	}

	res, err := ps.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &pr); err != nil {
		return
	}
	return
}

func (ps *PlansService) GetAll(page int) (pr *PlansResponse, err error) {
	ps.client.SetPage(page)
	req, err := ps.client.NewAPIRequest(true, http.MethodGet, "plans", nil)
	if err != nil {
		return nil, err
	}

	res, err := ps.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &pr); err != nil {
		return
	}
	return
}

func (ps *PlansService) Get(token string) (pr *PlanResponse, err error) {
	u := fmt.Sprintf("plans/%s", token)
	req, err := ps.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := ps.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &pr); err != nil {
		return
	}
	return
}

func (ps *PlansService) Update(plan *Plan) (pr *PlanResponse, err error) {
	u := fmt.Sprintf("plans/%s", plan.Token)
	req, err := ps.client.NewAPIRequest(true, http.MethodPut, u, plan)
	if err != nil {
		return nil, err
	}

	res, err := ps.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &pr); err != nil {
		return
	}
	return
}

func (ps *PlansService) Delete(token string) (success bool, err error) {
	u := fmt.Sprintf("plans/%s", token)
	req, err := ps.client.NewAPIRequest(true, http.MethodDelete, u, nil)
	if err != nil {
		return false, err
	}

	res, err := ps.client.Do(req)
	if err != nil {
		return false, err
	}
	if res.StatusCode != 204 {
		return false, errors.New("plan not found")
	}

	return true, nil
}

func (ps *PlansService) CreatePlanSubscription(subscription *Subscription) (sr *SubscriptionResponse, err error) {
	u := fmt.Sprintf("plans/%s/subscriptions", subscription.PlanToken)
	req, err := ps.client.NewAPIRequest(true, http.MethodPost, u, subscription)
	if err != nil {
		return nil, err
	}

	res, err := ps.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &sr); err != nil {
		return
	}
	return
}

func (ps *PlansService) GetPlanSubscriptions(token string) (sr *SubscriptionsResponse, err error) {
	u := fmt.Sprintf("plans/%s/subscriptions", token)
	req, err := ps.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := ps.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &sr); err != nil {
		return
	}
	return
}
