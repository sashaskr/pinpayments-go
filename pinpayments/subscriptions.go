package pinpayments

import "time"

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

type SubscriptionsResponse struct {
	Response []Subscription `json:"response,omitempty"`
	Count      int                    `json:"count"`
	Pagination struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
		Next     int `json:"next"`
		PerPage  int `json:"per_page"`
		Pages    int `json:"pages"`
		Count    int `json:"count"`
	} `json:"pagination"`
}
