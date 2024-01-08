[![CodeFactor](https://www.codefactor.io/repository/github/rareplanet1/pinpayments-go/badge)](https://www.codefactor.io/repository/github/rareplanet1/pinpayments-go)
# Pin Payments
Golang client to the http://pinpayments.com/ API. 

Forked from [https://github.com/sashaskr/pinpayments-go][https://github.com/sashaskr/pinpayments-go], which appears unmaintained, and with the intention of fixing any outstanding issues and adding new functionality provided by Pin Payments since.

## How to use
First, you need to create a config:
```go
config := pinpayments.NewConfig(true, YOUR_SECRET_KEY_HERE, YOUR_PUBLISHABLE_KEY_HERE)
```
The first argument `false|true` is indicated `live|test` mode of your API

Then you need to create a client:
```go
client, err := pinpayments.NewClient(nil, config)
	if err != nil {
		panic(err) // Or your way of handling errors
	}
```
The first argument is the http client you are going to use. If set as `nil`, it means client is going to use default go http client. You can use your own

After the client is created, you can access to the services currently supported by this API:
```go
client.Charges
client.Customers
client.Refunds
client.Cards
client.Recipients
client.Transfers
client.Balance
client.BankAccount
client.Events
client.WebhookEndpoints
client.Webhooks
client.Plans
client.Subscription
client.Merchant //Currently BETA and you have to contact Pin Payments
```
Each of the services has methods, covered all declared API provied by Pin Payments. Each method returns the unmarshalled struct, so you can use it in your business logic

