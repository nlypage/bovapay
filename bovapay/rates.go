package bovapay

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Rates struct {
	UsdtRub float64 `json:"usdt_rub"`
	UsdtUah float64 `json:"usdt_uah"`
	UsdtUzs float64 `json:"usdt_uzs"`
	UsdtKgs float64 `json:"usdt_kgs"`
}

// GetRates is a function to get currency pair rates using merchant/rates endpoint.
func (c *Client) GetRates() (*Rates, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "merchant/rates",
	}

	resp, err := c.Do(r)
	if err != nil {
		return nil, err
	}

	var ratesResponse Rates
	if errUnmarshal := json.Unmarshal(resp, &ratesResponse); errUnmarshal != nil {
		return nil, fmt.Errorf("error while unmarshaling response to the target: %w", errUnmarshal)
	}

	return &ratesResponse, nil
}
