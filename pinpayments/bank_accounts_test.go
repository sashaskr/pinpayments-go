package pinpayments

import (
	"github.com/sashaskr/pinpayments-go/testdata"
	"net/http"
	"testing"
)

func TestBankAccountService_Create(t *testing.T) {
	setup()
	defer teardown()

	token := "ba_nytGw7koRg23EEp9NTmz9w"

	_ = tClient.WithAuthenticationValue("test_secret", "test_publishable")
	tMux.HandleFunc("/bank_accounts", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Basic dGVzdF9zZWNyZXQ6")
		testMethod(t, r, http.MethodPost)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.CreateBankAccountResponse))
	})

	bankAccount := BankAccount{
		Name:     "Mr Ronald Robot",
		Bsb:      "123456",
		Number:   "987654321",
	}

	res, err := tClient.BankAccount.Create(&bankAccount)
	if err != nil {
		t.Fatal(err)
	}

	if res.Response.Token != token {
		t.Errorf("mismatching info. want %s got %s",token, res.Response.Token)
	}

	if res.Response.Name != bankAccount.Name {
		t.Errorf("mismatching info. want %s got %s",bankAccount.Name, res.Response.Name)
	}
}
