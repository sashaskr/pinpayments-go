package helpers

const (
	VisaSuccessful       = "4200000000000000"
	MastercardSuccessful = "5520000000000000"
	AmexSuccessful       = "372000000000000"

	VisaDeclined       = "4100000000000001"
	MastercardDeclined = "5560000000000001"
	AmexDeclined       = "371000000000001"

	VisaInsufficientFunds       = "4000000000000002"
	MastercardInsufficientFunds = "5510000000000002"
	AmexInsufficientFunds       = "370000000000002"

	VisaInvalidCVV       = "4900000000000003"
	MastercardInvalidCVV = "5550000000000003"
	AmexInvalidCVV       = "379000000000003"

	VisaInvalidCard       = "4800000000000004"
	MastercardInvalidCard = "5500000000000004"
	AmexInvalidCard       = "378000000000004"

	VisaProcessingError       = "4700000000000005"
	MastercardProcessingError = "5590000000000005"
	AmexProcessingError       = "377000000000005"

	VisaSuspectedFraud       = "4600000000000006"
	MastercardSuspectedFraud = "5540000000000006"
	AmexSuspectedFraud       = "376000000000006"

	VisaGatewayError       = "4300000000000009"
	MastercardGatewayError = "5570000000000009"
	AmexGatewayError       = "373000000000009"

	VisaUnknownError       = "4400000000000099"
	MastercardUnknownError = "5530000000000099"
	AmexUnknownError       = "374000000000099"
)
