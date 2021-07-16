package pinpayments

import (
	"github.com/sashaskr/pinpayments-go/pinpayments/helpers"
	"github.com/sashaskr/pinpayments-go/testdata"
	"net/http"
	"testing"
)

func TestChargesService_CreateCharge(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_secret", "test_publishable")
	tMux.HandleFunc("/charges", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Basic dGVzdF9zZWNyZXQ6")
		testMethod(t, r, http.MethodPost)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.ChargeCreateResponse))
	})

	var chargeToken = "ch_lfUYEBK14zotCTykezJkfg"

	charge := ChargesRequest{
		Email:       "roland@pinpayments.com",
		Description: "test charge",
		Amount:      400,
		IpAddress:   "203.192.1.172",
		Currency:    "AUD",
		Capture:     true,
		Metadata: Metadata{
			OrderNumber:  "123456",
			CustomerName: "Roland Robot",
		},
		Card: Card{
			Token:           "card_pIQJKMs93GsCc9vLSLevbw",
			ExpiryMonth:     5,
			ExpiryYear:      2022,
			IssuingCountry:  "US",
			Number:          helpers.MastercardSuccessful,
			CVC:             "123",
			AddressLine1:    "42 Sevenoaks St",
			AddressCity:     "Lathlain",
			AddressPostcode: "6454",
			AddressState:    "WA",
			AddressCountry:  "Australia",
		},
	}
	res, err := tClient.Charges.CreateCharge(&charge)
	if err != nil {
		t.Fatal(err)
	}

	if res.Response.Token != chargeToken {
		t.Errorf("mismatching info. want %s got %s", chargeToken, res.Response.Token)
	}
}
