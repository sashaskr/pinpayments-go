package pinpayments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type WebhooksService service

type Webhook struct {
	Token                string      `json:"token,omitempty"`
	Status               string      `json:"status,omitempty"`
	Url                  url.URL     `json:"url,omitempty"`
	EventToken           string      `json:"event_token,omitempty"`
	WebhookEndpointToken string      `json:"webhook_endpoint_token,omitempty"`
	CreatedAt            time.Time   `json:"created_at,omitempty"`
	AcceptedAt           time.Time   `json:"accepted_at,omitempty"`
	NextRun              time.Time   `json:"next_run,omitempty"`
	Retries              int         `json:"retries,omitempty"`
	Errors               interface{} `json:"errors,omitempty"`
}

type WebhookResponse struct {
	Response Webhook `json:"response,omitempty"`
}

type WebhooksResponse struct {
	Response   []Webhook `json:"response,omitempty"`
	Pagination struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
		Next     int `json:"next"`
		PerPage  int `json:"per_page"`
		Pages    int `json:"pages"`
		Count    int `json:"count"`
	} `json:"pagination"`
}

func (ws *WebhooksService) GetAll(page int) (wr *WebhooksResponse, err error) {
	ws.client.SetPage(page)
	req, err := ws.client.NewAPIRequest(true, http.MethodGet, "webhooks", nil)
	if err != nil {
		return nil, err
	}

	res, err := ws.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &wr); err != nil {
		return
	}
	return
}

func (es *WebhooksService) Get(token string) (er *WebhookResponse, err error) {
	u := fmt.Sprintf("webhooks/%s", token)
	req, err := es.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := es.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &er); err != nil {
		return
	}
	return
}

func (es *WebhooksService) Replay(token string) (er *WebhookResponse, err error) {
	u := fmt.Sprintf("webhooks/%s/replay", token)
	req, err := es.client.NewAPIRequest(true, http.MethodPut, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := es.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &er); err != nil {
		return
	}
	return
}
