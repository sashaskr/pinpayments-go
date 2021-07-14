package pinpayments

import (
	"github.com/sashaskr/pinpayments-go/pinpayments/helpers"
	"github.com/sashaskr/pinpayments-go/testdata"
	"net/http"
	"testing"
	"time"
)

func TestCardsService_Create(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_secret", "test_publishable")
	tMux.HandleFunc("/cards", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Basic dGVzdF9zZWNyZXQ6")
		testMethod(t, r, http.MethodPost)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.CreateCardResponse))
	})

	cardExpirationDate := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	var cardToken = "card_pIQJKMs93GsCc9vLSLevbw"

	card := Card{
		Number:            helpers.MastercardSuccessful,
		ExpiryMonth:       12,
		ExpiryYear:        cardExpirationDate.Year(),
		CVC:               "123",
		Name:              "Roland Robot",
		AddressLine1:      "42 Sevenoaks St",
		AddressCity:       "Lathlain",
		AddressPostcode:   "6454",
		AddressState:      "WA",
		AddressCountry:    "Australia",
		DisplayNumber:     "XXXX-XXXX-XXXX-0000",
	}

	res, err := tClient.Cards.Create(&card)
	if err != nil {
		t.Fatal(err)
	}

	if res.Response.Token != cardToken {
		t.Errorf("mismatching info. want %s got %s",cardToken, res.Response.Token)
	}
	if res.Response.ExpiryYear != card.ExpiryYear {
		t.Errorf("mismatching info. want %d got %d",card.ExpiryYear, res.Response.ExpiryYear)
	}
	if res.Response.DisplayNumber != card.DisplayNumber {
		t.Errorf("mismatching info. want %s got %s",card.DisplayNumber, res.Response.DisplayNumber)
	}
}
