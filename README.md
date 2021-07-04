[![CodeFactor](https://www.codefactor.io/repository/github/sashaskr/pinpayments-go/badge)](https://www.codefactor.io/repository/github/sashaskr/pinpayments-go)
# PinPayment
Golang fully functional client to the http://pinpayments.com/ API.
## How to use
First, you need to create a config:
```go
config := pinpayments.NewConfig(true, YOUR_SECRET_KEY_HERE, YOUR_PUBLISHABLE_KEY_HERE)
```
The first argument `true|false` is indicated `live|test` mode of your API

Then you need to create a client:
```go
client, err := pinpayments.NewClient(nil, config)
	if err != nil {
		panic(err) // Or your way of handling errors
	}
```
The first argument is the http client you are going to use. If set as `nil`, it means client is going to use default go http client. You can use your own

Since the client is created, you can access to the services:
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
client.Merchant //Currently BETA and you have to contact with pinpayments
```
Each of the services has methods, covered all declared API provied by Pinpayments. Each method returns the unmarshalled struct, so you can use it in your business logic

You can ask me a questions via [https://github.com/sashaskr/pinpayments-go/discussions/categories/q-a][https://github.com/sashaskr/pinpayments-go/discussions/categories/q-a]

Enjoy!

[https://github.com/sashaskr/pinpayments-go/discussions/categories/q-a]: https://github.com/sashaskr/pinpayments-go/discussions/categories/q-a