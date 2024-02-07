package pinpayments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EventsService service

type Event struct {
	Token     string      `json:"token,omitempty"`
	Type      string      `json:"type,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
}

type EventResponse struct {
	Response Event `json:"response,omitempty"`
}

type EventsResponse struct {
	PaginationInterface
	Response   []Event `json:"response,omitempty"`
	Count      int     `json:"count"`
	Pagination struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
		Next     int `json:"next"`
		PerPage  int `json:"per_page"`
		Pages    int `json:"pages"`
		Count    int `json:"count"`
	} `json:"pagination"`
}

func (es *EventsService) GetAll(page int) (er *EventsResponse, err error) {
	es.client.SetPage(page)
	req, err := es.client.NewAPIRequest(true, http.MethodGet, "events", nil)

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

func (es *EventsService) Get(token string) (er *EventResponse, err error) {
	u := fmt.Sprintf("events/%s", token)
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
