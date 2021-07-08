package pinpayments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WebhookEndpointsService service

type Endpoint struct {
	Token     string    `json:"token,omitempty"`
	Key       string    `json:"key,omitempty"`
	Url       string    `json:"url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type EndpointResponse struct {
	Response Endpoint `json:"response,omitempty"`
}

type EndpointsResponse struct {
	Response   []Endpoint `json:"response,omitempty"`
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

func (es *WebhookEndpointsService) Create(endpoint *Endpoint) (er *EndpointResponse, err error) {
	req, err := es.client.NewAPIRequest(true, http.MethodPost, "webhook_endpoints", endpoint)
	if err != nil {
		panic(err)
	}

	res, err := es.client.Do(req)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(res.content, &er); err != nil {
		return
	}
	return
}

func (es *WebhookEndpointsService) GetAll(page int) (er *EndpointsResponse, err error) {
	es.client.SetPage(page)
	req, err := es.client.NewAPIRequest(true, http.MethodGet, "webhook_endpoints", nil)
	if err != nil {
		panic(err)
	}

	res, err := es.client.Do(req)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(res.content, &er); err != nil {
		return
	}
	return
}

func (es *WebhookEndpointsService) Get(token string) (er *EndpointResponse, err error) {
	u := fmt.Sprintf("webhook_endpoints/%s", token)
	req, err := es.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		panic(err)
	}

	res, err := es.client.Do(req)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(res.content, &er); err != nil {
		return
	}
	return
}

func (es *WebhookEndpointsService) Delete(token string) (er bool, err error) {
	u := fmt.Sprintf("webhook_endpoints/%s", token)
	req, err := es.client.NewAPIRequest(true, http.MethodDelete, u, nil)
	if err != nil {
		panic(err)
	}

	res, err := es.client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 204 {
		panic("user not found")
	}

	return true, nil
}
