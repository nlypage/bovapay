package bovapay

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateDepositRequest struct {
	// Your transaction number is on its side.
	MerchantID string `json:"merchant_id"`

	// The amount of the application in the format "0.00".
	Amount float64 `json:"amount"`

	// User ID who will pay for the application as a string in any format.
	PayeerIdentifier string `json:"payeer_identifier"`

	//IP of the user who will pay for the application.
	PayeerIP string `json:"payeer_ip"`

	// Primary or Secondary is indicated: "ftd" or "trust".
	PayeerType payeerType `json:"payeer_type"`

	// Application lifetime specified in seconds (up to a maximum of 30 minutes)
	Lifetime int32 `json:"lifetime"`

	// When choosing crypto_currency as a merchant, the transaction will not be empty, it will be processed by a crypto transaction
	Currency currency `json:"currency"`

	// *Optional
	// You can create an empty request, where the user will select the payment method available to him.
	// You can create a request with the selected payment method by the merchant by specifying the payment method in the request.
	PaymentMethod paymentMethod `json:"payment_method"`

	// Your webhook URL
	CallbackURL string `json:"callback_url"`

	// *Optional
	// The URL to which the user will be redirected after payment.
	RedirectURL string `json:"redirect_url"`

	// *Optional
	// Email of the user who will pay for the application.
	Email string `json:"email"`

	// *Optional
	// Customer name who will pay for the application.
	CustomerName string `json:"customer_name"`
}

type Deposit struct {
	UUID                   string  `json:"uuid"`
	MerchantID             string  `json:"merchant_id"`
	Amount                 float64 `json:"amount,string"`
	FiatAmount             float64 `json:"fiat_amount,string"`
	Currency               string  `json:"currency"`
	State                  string  `json:"state"`
	SelectedCryptoCurrency string  `json:"selected_crypto_currency"`
	CallbackURL            string  `json:"callback_url"`
	RedirectURL            string  `json:"redirect_url"`
	CreatedAt              string  `json:"created_at"`
	UpdatedAt              string  `json:"updated_at"`
	FormURL                string  `json:"form_url"`
	SourceTransactionClass string  `json:"source_transaction_class"`
	//SourceTransaction any `json:"source_transaction"`
}

type CreateDepositResponse struct {
	merchantResponse
	Data Deposit `json:"data"`
}

// CreateDeposit creates a new deposit using the merchant/v1/deposits endpoint
func (c *Client) CreateDeposit(createDepositRequest CreateDepositRequest) (*Deposit, error) {
	r := &request{
		method:            http.MethodPost,
		endpoint:          "merchant/v1/deposits",
		authorizationType: signatureAuthorization,
	}
	r.Add("user_uuid", c.userUuid)
	r.Add("merchant_id", createDepositRequest.MerchantID)
	r.Add("amount", createDepositRequest.Amount)
	r.Add("payeer_identifier", createDepositRequest.PayeerIdentifier)
	r.Add("payeer_ip", createDepositRequest.PayeerIP)
	r.Add("payeer_type", createDepositRequest.PayeerType)
	r.Add("lifetime", createDepositRequest.Lifetime)
	r.Add("currency", createDepositRequest.Currency)
	if createDepositRequest.PaymentMethod != "" {
		r.Add("payment_method", createDepositRequest.PaymentMethod)
	}
	if createDepositRequest.CallbackURL != "" {
		r.Add("callback_url", createDepositRequest.CallbackURL)
	}
	if createDepositRequest.RedirectURL != "" {
		r.Add("redirect_url", createDepositRequest.RedirectURL)
	}
	if createDepositRequest.Email != "" {
		r.Add("email", createDepositRequest.Email)
	}
	if createDepositRequest.CustomerName != "" {
		r.Add("customer_name", createDepositRequest.CustomerName)
	}

	resp, err := c.Do(r)
	if err != nil {
		return nil, err
	}

	var createResponse CreateDepositResponse
	if errUnmarshal := json.Unmarshal(resp, &createResponse); errUnmarshal != nil {
		return nil, fmt.Errorf("error while unmarshaling response to the target: %w", errUnmarshal)
	}

	if createResponse.Status == "ok" {
		return &createResponse.Data, nil
	} else {
		return nil, fmt.Errorf("createDeposit request error: status - %s, errors - %v, message - %s", createResponse.Status, createResponse.Errors, *createResponse.Message)
	}
}
