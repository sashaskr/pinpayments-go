package testdata

const CreateBankAccountRequest = `{
	"name": "Mr Ronald Robot",
	"bsb": "123456",
	"number": "987654321"
}`

const CreateBankAccountResponse = `
{
  "response": {
    "token": "ba_nytGw7koRg23EEp9NTmz9w",
    "name": "Mr Ronald Robot",
    "bsb": "123456",
    "number": "XXXXXX321",
    "bank_name": "",
    "branch": ""
  },
  "ip_address": "127.0.0.1"
}
`