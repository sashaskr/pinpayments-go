package pinpayments

import (
	"github.com/sashaskr/pinpayments-go/testdata"
	"net/http"
	"testing"
)

func TestBalanceService_GetBalance(t *testing.T) {
	setup()
	defer teardown()

	balanceAvailableAmount := 400
	balanceAvailableCurrency := "AUD"
	balancePendingAmount := 1200
	balancePendingCurrency := "NZD"

	_ = tClient.WithAuthenticationValue("test_secret", "test_publishable")
	tMux.HandleFunc("/balance", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Basic dGVzdF9zZWNyZXQ6")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetBalanceResponse))
	})

	res, err := tClient.Balance.GetBalance()
	if err != nil {
		t.Fatal(err)
	}

	if res.Response.Available[0].Amount != balanceAvailableAmount {
		t.Errorf("mismatching info. want %v, got %v", balanceAvailableAmount, res.Response.Available[0].Amount)
	}

	if res.Response.Available[0].Currency != balanceAvailableCurrency {
		t.Errorf("mismatching info. want %v, got %v", balanceAvailableCurrency, res.Response.Available[0].Currency)
	}

	if res.Response.Pending[0].Amount != balancePendingAmount {
		t.Errorf("mismatching info. want %v, got %v", balancePendingAmount, res.Response.Pending[0].Amount)
	}

	if res.Response.Pending[0].Currency != balancePendingCurrency {
		t.Errorf("mismatching info. want %v, got %v", balancePendingCurrency, res.Response.Pending[0].Currency)
	}
}