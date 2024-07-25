package bovapay

type currency string
type payeerType string
type paymentMethod string

const (
	RUB            = currency("rub")
	CryptoCurrency = currency("crypto_currency")
)

const (
	Primary   = payeerType("ftd")
	Secondary = payeerType("trust")
)

const (
	Card          = paymentMethod("card")
	SberPay       = paymentMethod("sberpay")
	International = paymentMethod("international")
)

type response struct {
	Payload    any    `json:"payload"`
	ResultCode string `json:"result_code"`
}

type merchantResponse struct {
	Data    any                    `json:"data"`
	Message *string                `json:"message"`
	Errors  map[string]interface{} `json:"errors"`
	Status  string                 `json:"status"`
	Meta    map[string]interface{} `json:"meta"`
}

// Transaction states
const (
	StatusCreated                   string = "created"
	StatusProcessing                string = "processing"
	StatusWaitingPayment            string = "waiting_payment"
	StatusPaid                      string = "paid"
	StatusSuccessed                 string = "successed"
	StatusFailed                    string = "failed"
	StatusClosedFailed              string = "closed_failed"
	StatusRepeatedClosedFailed      string = "repeated_closed_failed"
	StatusAcceptedSuccessed         string = "accepted_successed"
	StatusRepeatedAcceptedSuccessed string = "repeated_accepted_successed"
)
