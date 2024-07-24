package bovapay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nlypage/bovapay/bovapay/common"
	"io"
	"net/http"
	"time"
)

const (
	apiUrl = "https://bovatech.cc/"
)

type Client struct {
	apiKey   string
	userUuid string

	httpClient *http.Client
}

type Options struct {
	// APIKey from your bovapay personal cabinet.
	APIKey string
	// UserID from your bovapay personal cabinet.
	UserID string

	// ClientTimeout field is optional. Default is 30s.
	ClientTimeout time.Duration
}

// NewClient creates a new bovapay client to interact with api.
func NewClient(options Options) *Client {
	c := &Client{
		apiKey:   options.APIKey,
		userUuid: options.UserID,
	}
	clientTimeout := 30 * time.Second
	if options.ClientTimeout != 0 {
		clientTimeout = options.ClientTimeout
	}

	c.httpClient = &http.Client{
		Timeout: clientTimeout,
	}
	return c
}

// CompareSignature compares the signature from the webhook request with the expected signature.
func (c *Client) CompareSignature(body []byte, headers map[string][]string) bool {
	expectedSignature := common.GenerateSignature(string(body), c.apiKey)

	signature := ""
	if sigs, ok := headers["Signature"]; ok && len(sigs) > 0 {
		signature = sigs[0]
	} else {
		return false
	}

	return expectedSignature == signature
}

// Do send a request to the bovapay api.
func (c *Client) Do(method, endpoint string, body map[string]interface{}) ([]byte, error) {
	url := apiUrl + endpoint

	jsonBody, errMarshal := json.Marshal(body)
	if errMarshal != nil {
		return nil, errMarshal
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	signature := common.GenerateSignature(string(jsonBody), c.apiKey)

	req.Header.Set("Signature", signature)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error response from the server: %s", responseBody)
	}

	return responseBody, nil
}
