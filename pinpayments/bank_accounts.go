package pinpayments

import (
	"encoding/json"
	"net"
	"net/http"
)

type BankAccountService service

type BankAccount struct {
	Token    string `json:"token,omitempty"`
	Name     string `json:"name,omitempty"`
	Bsb      string `json:"bsb,omitempty"`
	Number   string `json:"number,omitempty"`
	BankName string `json:"bank_name,omitempty"`
	Branch   string `json:"branch,omitempty"`
}

type BankAccountResponse struct {
	Response  BankAccount `json:"response,omitempty"`
	IpAddress net.IP      `json:"ip_address,omitempty"`
}

func (bas *BankAccountService) Create(bankAccount *BankAccount) (bar *BankAccountResponse, err error) {
	req, err := bas.client.NewAPIRequest(true, http.MethodPost, "bank_accounts", bankAccount)
	if err != nil {
		return nil, err
	}

	res, err := bas.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.content, &bar); err != nil {
		return
	}
	return
}
