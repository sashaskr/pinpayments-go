package pinpayments

type BankAccount struct {
	Token    string `json:"token,omitempty"`
	Name     string `json:"name,omitempty"`
	Bsb      string `json:"bsb,omitempty"`
	Number   string `json:"number,omitempty"`
	BankName string `json:"bank_name,omitempty"`
	Branch   string `json:"branch,omitempty"`
}
