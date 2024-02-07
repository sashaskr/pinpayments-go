package pinpayments

import (
	"encoding/json"
	"net"
	"net/http"
)

type CardsService service

type Card struct {
	Number            string `json:"number,omitempty"`
	ExpiryMonth       int    `json:"expiry_month,omitempty"`
	ExpiryYear        int    `json:"expiry_year,omitempty"`
	CVC               string `json:"cvc,omitempty"`
	Name              string `json:"name,omitempty"`
	AddressLine1      string `json:"address_line1,omitempty"`
	AddressLine2      string `json:"address_line2,omitempty"`
	AddressCity       string `json:"address_city,omitempty"`
	AddressPostcode   string `json:"address_postcode,omitempty"`
	AddressState      string `json:"address_state,omitempty"`
	AddressCountry    string `json:"address_country,omitempty"`
	Scheme            string `json:"scheme,omitempty"`
	DisplayNumber     string `json:"display_number,omitempty"`
	IssuingCountry    string `json:"issuing_country,omitempty"`
	CustomerToken     string `json:"customer_token,omitempty"`
	Token             string `json:"token,omitempty"`
	PublishableApiKey string `json:"publishable_api_key,omitempty"`
}

type CardCreatedResponse struct {
	Response  Card   `json:"response,omitempty"`
	IpAddress net.IP `json:"ip_address,omitempty"`
}

func (cs *CardsService) Create(card *Card) (cr *CardCreatedResponse, err error) {
	req, err := cs.client.NewAPIRequest(true, http.MethodPost, "cards", card)
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
