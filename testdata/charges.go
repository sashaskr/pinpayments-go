package testdata

const ChargeCreateResponse = `{
	 "response": {
    "token": "ch_lfUYEBK14zotCTykezJkfg",
    "success": true,
    "amount": 400,
    "currency": "AUD",
    "description": "test charge",
    "email": "roland@pinpayments.com",
    "ip_address": "203.192.1.172",
    "created_at": "2012-06-20T03:10:49Z",
    "status_message": "Success",
    "error_message": null,
    "card": {
      "token": "card_pIQJKMs93GsCc9vLSLevbw",
      "scheme": "master",
      "display_number": "XXXX-XXXX-XXXX-0000",
      "issuing_country": "US",
      "expiry_month": 5,
      "expiry_year": 2022,
      "name": "Roland Robot",
      "address_line1": "42 Sevenoaks St",
      "address_line2": null,
      "address_city": "Lathlain",
      "address_postcode": "6454",
      "address_state": "WA",
      "address_country": "Australia",
      "customer_token": null,
      "primary": null
    },
    "transfer": [],
    "amount_refunded": 0,
    "total_fees": 42,
    "merchant_entitlement": 358,
    "refund_pending": false,
    "authorisation_expired": false,
    "authorisation_voided": false,
    "captured": true,
    "captured_at": "2012-06-20T03:10:49Z",
    "settlement_currency": "AUD",
    "active_chargebacks": false,
    "metadata": {
      "OrderNumber": "123456",
      "CustomerName": "Roland Robot"
    }
  }
}`
