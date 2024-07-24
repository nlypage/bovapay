package bovapay

import (
	"encoding/json"
	"github.com/nlypage/bovapay/bovapay/common"
)

// CompareSignature compares the signature from the webhook request with the expected signature.
func CompareSignature(body []byte, headers map[string][]string, apiKey string) bool {
	expectedSignature := common.GenerateSignature(string(body), apiKey)

	signature := ""
	if sigs, ok := headers["Signature"]; ok && len(sigs) > 0 {
		signature = sigs[0]
	} else {
		return false
	}

	return expectedSignature == signature
}

// WebhookUpdate represents the data that you will receive from the webhook.
type WebhookUpdate struct {
	ID                  string  `json:"id"`
	MerchantID          string  `json:"merchant_id"`
	Status              string  `json:"status"`
	Message             string  `json:"message"`
	Currency            string  `json:"currency"`
	PaymentMethod       string  `json:"payment_method"`
	Rate                float64 `json:"rate"`
	Amount              float64 `json:"amount"`
	FiatAmount          float64 `json:"fiat_amount"`
	OldFiatAmount       float64 `json:"old_fiat_amount"`
	ServiceCommission   float64 `json:"service_commission"`
	PayeerCardNumber    *string `json:"payeer_card_number"`
	RecipientCardNumber *string `json:"recipient_card_number"`
}

// ParseWebhookUpdate parses the webhook update from the request body.
func ParseWebhookUpdate(data []byte) (*WebhookUpdate, error) {
	var webhookUpdate WebhookUpdate
	if err := json.Unmarshal(data, &webhookUpdate); err != nil {
		return nil, err
	}

	return &webhookUpdate, nil
}
