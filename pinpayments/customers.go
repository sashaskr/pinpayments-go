package pinpayments

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type CustomersService service

type CustomerRequest struct {
	Token            string `json:"token,omitempty"`
	Email            string `json:"email,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	PhoneNumber      string `json:"phone_number,omitempty"`
	Company          string `json:"company,omitempty"`
	Notes            string `json:"notes,omitempty"`
	Card             Card   `json:"card,omitempty"`
	CardToken        string `json:"card_token,omitempty"`
	PrimaryCardToken string `json:"primary_card_token,omitempty"`
}

type CustomerResponseBody struct {
	Token       string    `json:"token,omitempty"`
	Email       string    `json:"email,omitempty"`
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Notes       string    `json:"notes,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Card        Card      `json:"card,omitempty"`
}

type CustomerResponse struct {
	Response CustomerResponseBody `json:"response"`
}

type CustomersResponse struct {
	Responses  []CustomerResponseBody `json:"response"`
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

type CardsResponse struct {
	Response   []Card `json:"response,omitempty"`
	Count      int    `json:"count"`
	Pagination struct {
		Current  int `json:"current"`
		Previous int `json:"previous"`
		Next     int `json:"next"`
		PerPage  int `json:"per_page"`
		Pages    int `json:"pages"`
		Count    int `json:"count"`
	} `json:"pagination"`
}

type CardCreationResponse struct {
	Response  Card   `json:"response,omitempty"`
	IpAddress string `json:"ip_address,omitempty"`
}

type CardCreationRequest struct {
	Card          Card
	CustomerToken string
}

type CardDeletionRequest struct {
	CustomerToken string
	CardToken     string
}

func (cs *CustomersService) Create(customer *CustomerRequest) (cr *CustomerResponse, err error) {
	req, err := cs.client.NewAPIRequest(true, http.MethodPost, "customers", customer)
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

func (cs *CustomersService) GetAll(page int) (cr *CustomersResponse, err error) {
	cs.client.SetPage(page)
	req, err := cs.client.NewAPIRequest(true, http.MethodGet, "customers", nil)
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

func (cs *CustomersService) Get(token string) (cr *CustomerResponse, err error) {
	u := fmt.Sprintf("customers/%s", token)
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

func (cs *CustomersService) Update(customerUpdated *CustomerRequest) (cr *CustomerResponse, err error) {
	u := fmt.Sprintf("customers/%s", customerUpdated.Token)
	req, err := cs.client.NewAPIRequest(true, http.MethodPut, u, customerUpdated)
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

func (cs *CustomersService) Delete(token string) (success bool, err error) {
	u := fmt.Sprintf("customers/%s", token)
	req, err := cs.client.NewAPIRequest(true, http.MethodDelete, u, nil)
	if err != nil {
		return false, err
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return false, err
	}
	if res.StatusCode != 204 {
		return false, errors.New("customer not found")
	}

	return true, nil
}

func (cs *CustomersService) GetCharges(token string) (cr *ChargesResponse, err error) {
	u := fmt.Sprintf("customers/%s/charges", token)
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

func (cs *CustomersService) GetCards(token string) (cr *CardsResponse, err error) {
	u := fmt.Sprintf("customers/%s/cards", token)
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

func (cs *CustomersService) AddCard(cardRequest *CardCreationRequest) (cr *CardCreationResponse, err error) {
	u := fmt.Sprintf("customers/%s/cards", cardRequest.CustomerToken)
	req, err := cs.client.NewAPIRequest(true, http.MethodPost, u, cardRequest.Card)
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

func (cs *CustomersService) DeleteCard(cardDeletion *CardDeletionRequest) (bool bool, err error) {
	u := fmt.Sprintf("customers/%s/cards/%s", cardDeletion.CustomerToken, cardDeletion.CardToken)
	req, err := cs.client.NewAPIRequest(true, http.MethodDelete, u, nil)
	if err != nil {
		return false, err
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return false, err
	}
	if res.StatusCode != 204 {
		return false, errors.New("card not found")
	}

	return true, nil
}

func (cs *CustomersService) GetSubscriptions(token string) (csub *SubscriptionsResponse, err error) {
	u := fmt.Sprintf("customers/%s/subscriptions", token)
	req, err := cs.client.NewAPIRequest(true, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &csub); err != nil {
		return
	}
	return
}
