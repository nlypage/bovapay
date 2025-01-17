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

// Do send a request to the bovapay api.
func (c *Client) Do(r *request) ([]byte, error) {
	url := apiUrl + r.endpoint

	jsonBody, errMarshal := json.Marshal(r.body)
	if errMarshal != nil {
		return nil, errMarshal
	}

	req, err := http.NewRequest(r.method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	switch r.authorizationType {
	case signatureAuthorization:
		signature := common.GenerateSignature(string(jsonBody), c.apiKey)

		req.Header.Set("Signature", signature)

	case authorizationToken:
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}

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
