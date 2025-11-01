package pinpayments

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type BalanceService service

type Balance struct {
	Amount   int    `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type Balances struct {
	Available []Balance `json:"available,omitempty"`
	Pending   []Balance `json:"pending,omitempty"`
}

type BalancesResponse struct {
	Response Balances `json:"response,omitempty"`
}

func (bc *BalanceService) GetBalance() (br *BalancesResponse, err error) {
	req, err := bc.client.NewAPIRequest(true, http.MethodGet, "balance", nil)
	if err != nil {
		return nil, errors.Wrap(err, "GetBalance:NewApiRequest")
	}

	res, err := bc.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "GetBalance:Do")
	}

	if err = json.Unmarshal(res.content, &br); err != nil {
		return nil, errors.Wrap(err, "GetBalance:Unmarshal")
	}
	return
}
